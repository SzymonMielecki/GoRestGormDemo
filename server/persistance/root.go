package persistance

import (
	"fmt"

	"github.com/SzymonMielecki/ksiazki/server/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}


func NewDB(host, user, password, dbname, port string) (*DB, error) {
	dsn := "host="+host+" user="+user+" password="+password+" dbname="+dbname+" port="+port +" sslmode=disable TimeZone=Europe/Warsaw"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&types.BookModel{}, &types.AuthorModel{}, &types.GenreModel{}); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db DB) GetBooks() ([]types.BookModel, error) {
	var books []types.BookModel
	err := db.Model(&types.BookModel{}).Preload("Author").Preload("Genre").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (db DB) GetBook(id string) (*types.BookModel, error) {
	var book types.BookModel
	err := db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (db DB) CreateBook(book *types.Book) error {
	author := db.GetOrCreateAuthor(book.Author)
	genre := db.GetOrCreateGenre(book.Genre)
	bookModel := types.NewBookModel(book, author, genre)
	
	if err := db.Create(&bookModel).Error; err != nil {
		fmt.Println("failed to create book: ", err)
		return err
	}
	if err := db.Model(&bookModel).Association("Author").Append(&author); err != nil {
		fmt.Println("failed to append author: ", err)
		return err
	}
	if err := db.Model(&bookModel).Association("Genre").Append(&genre); err != nil {
		fmt.Println("failed to append genre: ", err)
		return err
	}

	return nil
}

func (db DB) CreateAuthor(authorName string) error {
	return db.Create(types.NewAuthorModel(authorName)).Error
}

func (db DB) GetOrCreateAuthor(authorName string) types.AuthorModel {
	author := types.AuthorModel{Name: authorName}
	db.Where(types.AuthorModel{Name: authorName}).FirstOrCreate(&author)
	return author
}

func (db DB) CreateGenre(genreName string) error {
	return db.Create(types.NewGenreModel(genreName)).Error
}

func (db DB) GetOrCreateGenre(genreName string) types.GenreModel {
	genre := types.GenreModel{}
	db.Where(types.GenreModel{Name: genreName}).FirstOrCreate(&genre)
	return genre
}

func (db DB) Drop() error {
	if err := db.Migrator().DropTable(&types.BookModel{}, &types.AuthorModel{}, &types.GenreModel{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&types.BookModel{}, &types.AuthorModel{}, &types.GenreModel{}); err != nil {
		return  err
	}
	return nil
}