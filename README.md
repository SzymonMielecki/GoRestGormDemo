# Librarian - GoRestGormDemo

Librarian is a CLI for managing books. It is a simple application that allows you to add books to a database and search for them.

## Technologies used

-   [x] Go
-   [x] Docker
-   [x] Postgres
-   [x] Gorm
-   [x] Echo

## Usage Details

### Run the application

```bash
docker-compose up --build
```

### Run the Client

```bash
go build -o librarian client/main.go
./librarian
```

```
Librarian is a CLI for managing books

Usage:
  librarian [flags]

Flags:
  -a, --author string   Author of the book
  -g, --genre string    Genre of the book
  -h, --help            help for librarian
  -t, --title string    Title of the book
```
