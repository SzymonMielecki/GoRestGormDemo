package types

import "gorm.io/gorm"

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
		Url:    b.URL,
		Genre:  genre,
		Author: author,
	}
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Url   string `json:"url"`
	Genre  string `json:"genre"`
}

type BookModel struct {
	gorm.Model
	Title  string 
	Author AuthorModel 	
	AuthorID uint
	Url   string 
	Genre  GenreModel 
	GenreID uint
} 

type AuthorModel struct {
	gorm.Model
	Name string 
}

type GenreModel struct {
	gorm.Model
	Name string 
}

func NewBookModel(book *Book, author AuthorModel, genre GenreModel) *BookModel {
	return &BookModel{
		Title: book.Title,
		Author: author,
		AuthorID: author.ID,
		Url: book.Url,
		Genre: genre,
		GenreID: genre.ID,
	}
}

func (b *BookModel) ToBook() *Book {
	return &Book{
		Title: b.Title,
		Author: b.Author.Name,
		Url: b.Url,
		Genre: b.Genre.Name,
	}
}

func NewAuthorModel(author string) *AuthorModel {
	return &AuthorModel{Name: author}
}

func NewGenreModel(genre string) *GenreModel {
	return &GenreModel{Name: genre}
}