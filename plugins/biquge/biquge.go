package biquge

import "app-novel-downloader/engine"

type downloader struct {
    rootUrl string
}

func (dl *downloader) FetchBook(id string) *engine.Book {
    return nil
}

func (dl *downloader) FetchBookItem(bi *engine.BookItem) {

}

var Instance = &downloader{
    rootUrl: "https://www.biqugen.net",
}
