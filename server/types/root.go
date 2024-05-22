package types

import "gorm.io/gorm"

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
