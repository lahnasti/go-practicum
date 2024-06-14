package main

import ("fmt"
	"strings"
	"bufio"
	"os"
)

func main () {
	//сканер для чтения ввода
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите текст: ")
	scanner.Scan()
	text := scanner.Text()

	//разбитие текста на слова
	words := strings.Fields(text)
	// убрать знаки препинания

	wordCount := make(map[string]int)

	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	fmt.Println(wordCount)
}