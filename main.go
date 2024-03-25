package main

import (
    "app-novel-downloader/biqg"
    "app-novel-downloader/engine"
    "fmt"
)

func main() {
    dl := engine.NewDownloader(10, biqg.FetchBook, biqg.FetchBookItem)
    err := dl.StoreNovel("3979")
    if err != nil {
        fmt.Println("下载小说失败", err)
        return
    }
}
