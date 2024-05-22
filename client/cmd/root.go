package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
		return url + "books/" + strings.ToLower(strings.ReplaceAll(title, " ", "-")) +  "/", nil
	}
	if author != "" {
		url = url + "authors/" + strings.ToLower(strings.ReplaceAll(author, " ", "-")) +  "/"
	}
	if genre != "" {
		url = url + "genres/" + strings.ToLower(strings.ReplaceAll(genre, " ", "-")) +  "/"
	}
	return url+ "books", nil
}

var rootCmd = &cobra.Command{
	Use:   "air_qual",
	Short: "Check air quality in your area",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		url, err := getUrl()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println("URL: ", url)

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
		books := new([]types.Book)
		err = json.Unmarshal(body,books )
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		for book := range *books {
			fmt.Println("Title: ", (*books)[book].Title)
			fmt.Println("Author: ", (*books)[book].Author)
			fmt.Println("Genre: ", (*books)[book].Genre)
			fmt.Println("Href: ", (*books)[book].Href)
			fmt.Println()
		}

		// port := os.Getenv("PORT")
		// if port == "" {
		// 	port = ":8080"
		// }
		//
		//
		//
		// contentType := "application/json"
		//
		// if _, err = http.Post(url_weather, contentType, bytes.NewBuffer(body)); err != nil {
		// 	fmt.Println("Error: ", err)
		// 	return
		// }
		// fmt.Println("Data sent to server")
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
