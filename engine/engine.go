package engine

import (
    "fmt"
    "os"
    "strconv"
)

type Downloader struct {
    GoroutineNums   int
    Chan            chan *DLTask
    FetchBook       func(string) *Book
    FetchBookItem   func(*BookItem)
}

func (dl *Downloader) Download(id string) (*Book, error) {
    book := dl.FetchBook(id)
    if book.Error != nil {
        book.Done = true
        return book, book.Error
    }
    tasks := make([]*DLTask, 0, len(book.Items))
    dl.Chan = make(chan *DLTask, len(book.Items))
    for _, bi := range book.Items {
        t := NewDLTask(bi)
        dl.Chan <- t
        tasks = append(tasks, t)
    }
    close(dl.Chan)
    for i := 0; i < dl.GoroutineNums; i++ {
        go func() {
            for {
                t, ok := <- dl.Chan;
                if !ok {
                    break
                }
                dl.FetchBookItem(t.BookItem)
                t.Chan <- true
                close(t.Chan)
            }
        }()
    }
    for _, t := range tasks {
        <- t.Chan
        bi := t.BookItem
        if bi.Error != nil {
            fmt.Printf("下载章节处理错误,章节信息:Id=%d,Url=%s,Name=%s,错误信息:%v\n", bi.Id, bi.Url, bi.Name, bi.Error)
            continue
        }
        fmt.Printf("下载章节处理成功,章节信息:Id=%d,Url=%s,Name=%s\n", bi.Id, bi.Url, bi.Name)
    }
    book.Done = true
    return book, nil
}

func (dl *Downloader) StoreNovel(id string) error {
    book, err := dl.Download(id)
    if err != nil {
        return err
    }
    file, err := os.OpenFile("./" + book.Name + ".txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        return err
    }
    fmt.Println("下载小说章节总数:" + strconv.Itoa(len(book.Items)))
    for _, bi := range book.Items {
        sc := "\n\n" + bi.Name + "\n"
        for _, s := range bi.Sections {
            sc += s + "\n"
        }
        _, err = file.WriteString(sc)
        if err != nil {
            fmt.Printf("存储章节处理错误,章节信息:Id=%d,Url=%s,Name=%s,错误信息:%v\n", bi.Id, bi.Url, bi.Name, bi.Error)
            return err
        }
        fmt.Printf("存储章节处理成功,章节信息:Id=%d,Url=%s,Name=%s\n", bi.Id, bi.Url, bi.Name)
    }
    fmt.Println("存储小说章节总数:" + strconv.Itoa(len(book.Items)))
    fmt.Println("存储小说文件名:" + file.Name())
    return nil
}

func NewDownloader(goroutineNums int, fetchBook func(string) *Book, fetchBookItem func(*BookItem)) *Downloader {
    if goroutineNums <= 0 || goroutineNums > 50 {
        goroutineNums = 10
    }
    return &Downloader{
        GoroutineNums: goroutineNums,
        Chan: nil,
        FetchBook: fetchBook,
        FetchBookItem: fetchBookItem,
    }
}

type DLTask struct {
    BookItem        *BookItem
    Chan            chan bool
}

func NewDLTask(bi *BookItem) *DLTask {
    return &DLTask{
        BookItem: bi,
        Chan: make(chan bool, 1),
    }
}
