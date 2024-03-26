package plugins

import (
    "app-novel-downloader/plugins/biqg"
    "app-novel-downloader/plugins/bqgas"
)

var (
    // https://www.biqg.cc
    Biqg    = biqg.Instance
    // https://www.bqgas.cc
    Bqgas   = bqgas.Instance
)
