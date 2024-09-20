package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/SzymonMielecki/GoRestGormDemo/client/utils"
	"github.com/SzymonMielecki/GoRestGormDemo/types"
	"github.com/spf13/cobra"
)

var (
	title  string
	author string
	genre  string
)

var rootCmd = &cobra.Command{
	Use:   "librarian",
	Short: "Librarian is a CLI for managing books",
	Long:  "Librarian is a CLI for managing books",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		url, err := utils.GetUrl(title, author, genre)
		fmt.Println(url)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		booksPre := new([]types.BookPre)
		err = json.Unmarshal(body, booksPre)
		if err != nil {
			book := new(types.BookPre)
			err = json.Unmarshal(body, book)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			*booksPre = append(*booksPre, *book)
		}

		books := new([]types.Book)
		for _, bookPre := range *booksPre {
			*books = append(*books, *bookPre.ToBook())
		}

		backend := os.Getenv("BACKEND")
		if backend == "" {
			backend = "http://localhost:8080/books"
		}
		var wg sync.WaitGroup

		for _, book := range *books {
			wg.Add(1)
			go func(book types.Book) {
				defer wg.Done()
				body, err := json.Marshal(book)
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				res, err := http.Post(backend, "application/json", bytes.NewBuffer(body))
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				_ = res
				fmt.Println("Sent book: " + book.Author + " - \"" + book.Title + "\" to server with status:" + fmt.Sprint(res.StatusCode))
			}(book)
		}
		wg.Wait()
		fmt.Println("All data sent to server")
	},
}

func Connect() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the book")
	rootCmd.Flags().StringVarP(&author, "author", "a", "", "Author of the book")
	rootCmd.Flags().StringVarP(&genre, "genre", "g", "", "Genre of the book")
}
