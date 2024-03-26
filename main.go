package main

import (
    "app-novel-downloader/biqg"
    "app-novel-downloader/engine"
    "fmt"
)

func main() {
    err := engine.StoreNovel(biqg.Instance, "3979")
    if err != nil {
        fmt.Println("下载小说失败", err)
        return
    }
}
