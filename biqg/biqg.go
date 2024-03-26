package biqg

import (
    "app-novel-downloader/engine"
    "errors"
    "github.com/PuerkitoBio/goquery"
    "net/http"
    "strings"
)

const rootUrl = "https://www.biqg.cc"

func FetchBook(id string) *engine.Book {
    url := rootUrl + "/book/" + id
    book := engine.NewBook(url)
    book.Name = id
    resp, err := http.Get(url)
    if err != nil {
        book.Error = errors.New("获取章节列表页html,请求处理错误:" + err.Error())
        return book
    }
    defer func() { _ = resp.Body.Close() }()
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        book.Error = errors.New("解析章节列表页html,html解析错误:" + err.Error())
        return book
    }
    book.Html, _ = doc.Html()
    if title, exists := doc.Find(`html head meta[property="og:title"]`).Attr("content"); exists {
        book.Name = title
    }
    s := doc.Find(".listmain dl dd")
    if s == nil {
        book.Error = errors.New("解析章节列表页html,html元素缺失:章节列表元素nil")
        return book
    }
    s.Each(func(i int, s *goquery.Selection) {
        if s.HasClass("more pc_none") {
            return
        }
        a := s.ChildrenFiltered("a")
        href, _ := a.Attr("href")
        book.Items = append(book.Items, engine.NewBookItem(i, biqgUrl+ href, a.Text()))
    })
    return book
}

func FetchBookItem(bi *engine.BookItem) {
    resp, err := http.Get(bi.Url)
    if err != nil {
        bi.Error = errors.New("获取章节详情页html,请求处理错误:" + err.Error())
        return
    }
    defer func() { _ = resp.Body.Close() }()
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        bi.Error = errors.New("解析章节详情页html,html解析错误:" + err.Error())
        return
    }
    bi.Html, _ = doc.Html()
    content, _ := doc.Find("#chaptercontent").Html()
    bi.Sections = append(bi.Sections, strings.Split(content, "<br/><br/>")...)
    // 截取冗余内容(部分章节首行是章节名, 需移除)
    if len(bi.Sections) > 0 {
        if strings.Contains(bi.Name, strings.TrimSpace(bi.Sections[0])) {
            bi.Sections = bi.Sections[1:]
        }
    }
    // 截取冗余内容(章节末尾无效信息, 需移除)
    li := 0
    for i, text := range bi.Sections {
        if strings.Contains(text, rootUrl) {
            li = i
            break
        }
    }
    bi.Sections = bi.Sections[0:li]
}
