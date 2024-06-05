package endpoint

import (
	"net/http"

	"github.com/SzymonMielecki/GoDockerPsqlProject/server/logic"
	"github.com/SzymonMielecki/GoDockerPsqlProject/types"
	"github.com/labstack/echo/v4"
)

func GetBooks(a *logic.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		author := c.QueryParam("author")
		genre := c.QueryParam("genre")
		books, err := a.GetBooks()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if author != "" {
			books = a.FilterByAuthor(books, author)
		}
		if genre != "" {
			books = a.FilterByGenre(books, genre)
		}
		return c.JSON(http.StatusOK, books)
	}

}

func GetBook(a *logic.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		book, err := a.GetBook(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, book)
	}
}

func CreateBook(a *logic.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		book := new(types.Book)
		if err := c.Bind(book); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err := a.CreateBook(book)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusCreated, book)
	}
}

func Drop(a *logic.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := a.Drop()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, "Dropped")
	}
}
