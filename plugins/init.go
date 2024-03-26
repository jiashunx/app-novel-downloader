package plugins

import (
    "app-novel-downloader/plugins/biqg"
    "app-novel-downloader/plugins/biqooge"
    "app-novel-downloader/plugins/biquge"
    "app-novel-downloader/plugins/bqgas"
)

var (
    // https://www.biqg.cc
    Biqg    = biqg.Instance
    // https://www.biqooge.com
    Biqooge = biqooge.Instance
    // https://www.biqugen.net
    Biquge  = biquge.Instance
    // https://www.bqgas.cc
    Bqgas   = bqgas.Instance
)
