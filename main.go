package main

import (
    "app-novel-downloader/engine"
    "app-novel-downloader/plugins"
    "fmt"
)

func main() {
    err := engine.StoreNovel(plugins.Biqg, "3979")
    //err := engine.StoreNovel(plugins.Bqgas, "673")
    if err != nil {
        fmt.Println("下载小说失败", err)
        return
    }
}
