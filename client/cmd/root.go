package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/SzymonMielecki/ksiazki/client/utils"
	"github.com/SzymonMielecki/ksiazki/types"
	"github.com/spf13/cobra"
)

var (
	title string
	author string
	genre string
)


func getUrl() (string, error) {
	url := "https://wolnelektury.pl/api/"
	if title == "" && author == "" && genre == "" {
		return "", fmt.Errorf("you must provide at least one of the following flags: title, author, genre")
	}

	
	if title != "" {
		return url + "books/" + strings.ToLower(strings.ReplaceAll(utils.ReplacePolishChars(title), " ", "-")) +  "/", nil
	}
	if author != "" {
		url = url + "authors/" + strings.ToLower(strings.ReplaceAll(utils.ReplacePolishChars(author), " ", "-")) +  "/"
	}
	if genre != "" {
		url = url + "genres/" + strings.ToLower(strings.ReplaceAll(utils.ReplacePolishChars(genre), " ", "-")) +  "/"
	}
	return url+ "books", nil
}

var rootCmd = &cobra.Command{
	Use:   "ksiazki",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		url, err := getUrl()
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
		// fmt.Println( string(body))
		booksPre := new([]types.BookPre)
		err = json.Unmarshal(body,booksPre)
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
				res, err := http.Post(backend, "application/json",bytes.NewBuffer(body))
				if err != nil {	
					fmt.Println("Error: ", err)
					return
				}
				_ = res
				fmt.Println("Sent book: " + book.Author + " - \""+ book.Title+ "\" to server with status:"+ fmt.Sprint(res.StatusCode))
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
