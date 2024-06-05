package logic

import (
	"github.com/SzymonMielecki/GoDockerPsqlProject/server/persistance"
	"github.com/SzymonMielecki/GoDockerPsqlProject/types"
)

type AppState struct {
	db *persistance.DB
}

func NewAppState(db *persistance.DB) *AppState {
	return &AppState{db: db}
}

func (a *AppState) GetBooks() ([]types.Book, error) {
	books := new([]types.Book)
	bookModels, err := a.db.GetBooks()
	if err != nil {
		return nil, err
	}

	for _, bookModel := range bookModels {
		*books = append(*books, *bookModel.ToBook())
	}
	return *books, nil
}

func (a *AppState) GetBook(id string) (*types.Book, error) {
	bookModel, err := a.db.GetBook(id)
	if err != nil {
		return nil, err
	}
	return bookModel.ToBook(), nil
}

func (a *AppState) CreateBook(book *types.Book) error {
	return a.db.CreateBook(book)
}

func (a *AppState) Drop() error {
	return a.db.Drop()
}

func (a *AppState) FilterByAuthor(books []types.Book, author string) []types.Book {
	filteredBooks := []types.Book{}
	for _, book := range books {
		if book.Author == author {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks
}

func (a *AppState) FilterByGenre(books []types.Book, genre string) []types.Book {
	filteredBooks := []types.Book{}
	for _, book := range books {
		if book.Genre == genre {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks
}
