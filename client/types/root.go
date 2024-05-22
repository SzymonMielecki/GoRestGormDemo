package types

type BookPre struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Genres []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"genres"`
	Genre string `json:"genre"`
	Authors []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"authors"`
	Autor string `json:"author"`
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	URL   string `json:"url"`
	Genre  string `json:"genre"`
}

func (b *BookPre) ToBook() *Book {
	var author string
	if b.Autor != "" {
		author = b.Autor
	} else if len(b.Authors) > 0 {
		author = b.Authors[0].Name
	} else {
		author = "Unknown"
	}
	
	var genre string
	if b.Genre != "" {
		genre = b.Genre
	} else if len(b.Genres) > 0 {
		genre = b.Genres[0].Name
	} else  {
		genre = "Unknown"
	}

	return &Book{
		Title:  b.Title,
		URL:    b.URL,
		Genre:  genre,
		Author: author,
	}}
