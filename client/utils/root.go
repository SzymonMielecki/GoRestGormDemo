package utils

import "strings"

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