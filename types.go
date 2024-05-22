package types

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Href   string `json:"href"`
	Genre  string `json:"genre"`
}