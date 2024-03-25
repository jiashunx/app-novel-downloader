package engine

type Book struct {
    Url         string          `json:"url"`
    Name        string          `json:"name"`
    Items       []*BookItem     `json:"items"`
    Error       error           `json:"-"`
    Html        string          `json:"-"`
    Done        bool            `json:"-"`
}

func NewBook(url string) *Book {
    return &Book{
        Url: url,
        Name: "",
        Items: make([]*BookItem, 0, 0),
    }
}

type BookItem struct {
    Id          int             `json:"id"`
    Url         string          `json:"url"`
    Name        string          `json:"name"`
    Sections    []string        `json:"sections"`
    Error       error           `json:"-"`
    Html        string          `json:"-"`
}

func NewBookItem(id int, url, name string) *BookItem {
    return &BookItem{
        Id: id,
        Url: url,
        Name: name,
        Sections: make([]string, 0, 0),
    }
}
