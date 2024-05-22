package main

import (
	"log"
	"net/http"

	"github.com/SzymonMielecki/ksiazki/server/endpoint"
	"github.com/SzymonMielecki/ksiazki/server/logic"
	"github.com/SzymonMielecki/ksiazki/server/persistance"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	
	db,err := persistance.NewDB(
		"localhost",
		"postgres",
		"ksiazkiPass",
		"postgres",
		"5432",
	)
	if err != nil {
		log.Fatal(err)
	}
	a := logic.NewAppState(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker!")
	})
	e.GET("/books", endpoint.GetBooks(a))
	e.GET("/books/:id", endpoint.GetBook(a))
	e.POST("/books", endpoint.CreateBook(a))
	e.POST("/drop", endpoint.Drop(a))


	httpPort := "8080"

	e.Logger.Fatal(e.Start("localhost:" + httpPort))
}