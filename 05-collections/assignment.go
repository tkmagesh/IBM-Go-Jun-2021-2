package main

import (
	"fmt"
	"strings"
)

func main() {
	inputStr := "Aliquip magna irure eiusmod sit elit ipsum incididunt qui labore fugiat culpa ut commodo irure Excepteur duis Lorem dolor adipisicing dolor officia mollit voluptate Consectetur exercitation occaecat laboris voluptate amet aliqua esse ex anim enim Elit exercitation officia tempor consectetur consequat reprehenderit dolor labore ex veniam minim fugiat consequat Aliquip velit voluptate qui aute enim Magna labore sint nulla magna ipsum"
	words := strings.Split(inputStr, " ")
	wordsCountBySize := getWordsCountBySize(words)
	count, size := getMaxCountBySize(wordsCountBySize)
	fmt.Printf("The size of the word that occurs the most = %d and occurences = %d\n", size, count)
}

func getWordsCountBySize(words []string) map[int]int {
	wordsCountBySize := map[int]int{}
	for _, word := range words {
		wordsCountBySize[len(word)] += 1
	}
	return wordsCountBySize
}

func getMaxCountBySize(wordsBySize map[int]int) (count, size int) {
	for wordSize, wordCount := range wordsBySize {
		if wordCount > count {
			count = wordCount
			size = wordSize
		}
	}
	return
}
