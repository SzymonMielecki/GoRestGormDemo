package main

import (
	"fmt"
	"os"

	"github.com/SzymonMielecki/GoRestGormDemo/client/cmd"
)

func main() {
	Execute()
}

func Execute() {
	var title, author, genre string
	rootCmd := cmd.RootCommand(title, author, genre)
	rootCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the book")
	rootCmd.Flags().StringVarP(&author, "author", "a", "", "Author of the book")
	rootCmd.Flags().StringVarP(&genre, "genre", "g", "", "Genre of the book")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
