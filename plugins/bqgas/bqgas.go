package bqgas

import (
    "app-novel-downloader/engine"
    "app-novel-downloader/plugins/biqg"
)

type downloader struct {
    proxyDL engine.Downloader
}

func (dl *downloader) FetchBook(id string) *engine.Book {
    return dl.proxyDL.FetchBook(id)
}

func (dl *downloader) FetchBookItem(bi *engine.BookItem) {
    dl.proxyDL.FetchBookItem(bi)
}

var Instance = &downloader{
    proxyDL: biqg.NewDownloader("https://www.bqgas.cc"),
}
