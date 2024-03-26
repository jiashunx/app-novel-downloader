package main

import (
    "app-novel-downloader/engine"
    "app-novel-downloader/plugins"
    "fmt"
)

func main() {
    err := engine.StoreNovel(plugins.Biqg, "3979")
    if err != nil {
        fmt.Println("下载小说失败", err)
        return
    }
}
