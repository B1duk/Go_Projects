package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removePunctuation(str string) string {
	punctuation := []string{",", ".", "!", "?", ":", ";"}

	str = strings.ToLower(str)

	for _, v := range punctuation {
		str = strings.ReplaceAll(str, v, "")
	}
	return str
}

func splitIntoWords(str string) []string {
	words := strings.Fields(str)
	return words
}

func countWords(words []string) map[string]int {
	wordsMap := make(map[string]int)

	for _, v := range words {
		_, ok := wordsMap[v]
		if ok {
			wordsMap[v] += 1
		} else {
			wordsMap[v] = 1
		}
	}

	return wordsMap
}
func main() {
	fmt.Println("Введите строку")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Print("Ошибка ввода")
		return
	}
	cleanStr := removePunctuation(str)
	words := splitIntoWords(cleanStr)
	allWords := countWords(words)
	fmt.Println(allWords)
}
