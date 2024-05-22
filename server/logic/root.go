package logic

import (
	"github.com/SzymonMielecki/ksiazki/server/persistance"
	"github.com/SzymonMielecki/ksiazki/server/types"
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