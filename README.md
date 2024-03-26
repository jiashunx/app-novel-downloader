
### app-novel-downloader

客户端应用: 基于 Go 开发的小说下载器 (数据源: 笔趣阁 ), 可自行拓展

- 数据源

   - [https://www.biqg.cc](https://www.biqg.cc)

   - [https://www.bqgas.cc](https://www.bqgas.cc)

- 下载小说步骤

   - 首先确定要使用的数据源，例如：使用 [https://www.biqg.cc](https://www.biqg.cc)

   - 从数据源搜索要下载的小说，例如：《凡人修仙传》

   - 根据搜索结果跳转小说的章节目录页面，例如：凡人修仙传的章节目录页面为：[https://www.biqg.cc/book/3979/](https://www.biqg.cc/book/3979/)

   - 确定小说标识id，例如：凡人修仙传的小说标识id为：3979

   - 修改 [main.go](./main.go) 代码，选择相应plugins（下载器），并修改要下载的小说标识id

   - 下载后的小说以txt格式存储在 [main.go](./main.go) 文件同目录

- 拓展数据源步骤

   - 自定义plugin（实现方法 FetchBook 及 FetchBootItem，方法定义参见 [Downloader](./engine/engine.go) 接口）

   - 注册plugin（修改plugins包中的 [init.go](./plugins/init.go)，注册数据源下载器）

   - 修改 [main.go](./main.go) 代码，选择相应plugin（下载器），并从新plugin数据源中查询小说标识id
