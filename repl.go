package main
import "strings"

func cleanInput(text string) []string {
	lowStr := strings.ToLower(text)
	words := strings.Fields(lowStr)
	return words
}