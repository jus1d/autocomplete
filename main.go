package main

import (
	"autocomplete/trie"
	"fmt"
	"os"
	"strings"
)

func getDictionary(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		return []string{}
	}

	return strings.Split(string(content), "\n")
}

func main() {
	tree := trie.New()
	dict := getDictionary("./dictionary.txt")
	tree.InsertMany(dict)

	word := os.Args[1]

	results := tree.Autocomplete(word)

	fmt.Println("=========")
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
}
