package engine

import (
    "fmt"
    "os"
    "strconv"
)

type Downloader interface {
    FetchBook(id string) *Book
    FetchBookItem(bi *BookItem)
}

type task struct {
    bookItem *BookItem
    ch chan bool
}

func Download(dl Downloader, id string) (*Book, error) {
    book := dl.FetchBook(id)
    if book.Error != nil {
        book.Done = true
        return book, book.Error
    }
    fmt.Printf("获取小说章节列表成功,小说信息:Id=%s,Url=%s,Name=%s\n", id, book.Url, book.Name)
    tasks := make([]*task, 0, len(book.Items))
    ch := make(chan *task, len(book.Items))
    for _, bi := range book.Items {
        t := &task{
            bookItem: bi,
            ch: make(chan bool, 1),
        }
        ch <- t
        tasks = append(tasks, t)
    }
    close(ch)
    for i := 0; i < 10; i++ {
        go func() {
            for {
                t, ok := <- ch;
                if !ok {
                    break
                }
                dl.FetchBookItem(t.bookItem)
                t.ch <- true
                close(t.ch)
            }
        }()
    }
    for _, t := range tasks {
        <- t.ch
        bi := t.bookItem
        if bi.Error != nil {
            fmt.Printf("下载章节处理错误,章节信息:Id=%d,Url=%s,Name=%s,错误信息:%v\n", bi.Id, bi.Url, bi.Name, bi.Error)
            continue
        }
        fmt.Printf("下载章节处理成功,章节信息:Id=%d,Url=%s,Name=%s\n", bi.Id, bi.Url, bi.Name)
    }
    book.Done = true
    return book, nil
}

func StoreNovel(dl Downloader, id string) error {
    book, err := Download(dl, id)
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
