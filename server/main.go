package main

import (
	"fmt"
	"net/http"

	"github.com/SzymonMielecki/GoRestGormDemo/server/endpoint"
	"github.com/SzymonMielecki/GoRestGormDemo/server/logic"
	"github.com/SzymonMielecki/GoRestGormDemo/server/persistance"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	db, err := handleDbConnection()

	if err != nil {
		e.Logger.Fatal(err)
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

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}

func handleDbConnection() (*persistance.DB, error) {
	db, err := persistance.NewDB(
		"db",
		"postgres",
		"ksiazkiPass",
		"postgres",
		"5432",
	)
	if err == nil {
		return db, nil
	}
	db, err = persistance.NewDB(
		"localhost",
		"postgres",
		"ksiazkiPass",
		"postgres",
		"5432",
	)
	if err == nil {
		return db, nil
	}
	return nil, fmt.Errorf("could not connect to database")
}
