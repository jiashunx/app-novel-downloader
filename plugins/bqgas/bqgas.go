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
    // 与 https://www.biqg.cc 网页格式元素一致，因此复用 biqg 下载器
    proxyDL: biqg.NewDownloader("https://www.bqgas.cc"),
}
