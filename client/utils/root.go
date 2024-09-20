package utils

import (
	"fmt"
	"strings"
)

func ReplacePolishChars(input string) string {
	polishToASCII := map[rune]string{
		'ą': "a", 'ć': "c", 'ę': "e", 'ł': "l", 'ń': "n",
		'ó': "o", 'ś': "s", 'ź': "z", 'ż': "z",
		'Ą': "A", 'Ć': "C", 'Ę': "E", 'Ł': "L", 'Ń': "N",
		'Ó': "O", 'Ś': "S", 'Ź': "Z", 'Ż': "Z",
	}
	var builder strings.Builder

	for _, char := range input {
		if replacement, found := polishToASCII[char]; found {
			builder.WriteString(replacement)
		} else {
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

func GetUrl(title, author, genre string) (string, error) {
	url := "https://wolnelektury.pl/api/"
	if title == "" && author == "" && genre == "" {
		return "", fmt.Errorf("you must provide at least one of the following flags: title, author, genre")
	}

	if title != "" {
		return url + "books/" + strings.ToLower(strings.ReplaceAll(ReplacePolishChars(title), " ", "-")) + "/", nil
	}
	if author != "" {
		url = url + "authors/" + strings.ToLower(strings.ReplaceAll(ReplacePolishChars(author), " ", "-")) + "/"
	}
	if genre != "" {
		url = url + "genres/" + strings.ToLower(strings.ReplaceAll(ReplacePolishChars(genre), " ", "-")) + "/"
	}
	return url + "books", nil
}
