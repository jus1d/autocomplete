package main

import (
	"autocomplete/dict"
	"autocomplete/trie"
	"flag"
	"fmt"
	"os"
	"strings"
)

func getDictionary(path string) []string {
	if path == "" {
		return dict.Fruits
	}

	_, err := os.Stat(path)
	if err != nil {
		fmt.Printf("WARN: Couldn't found a dictionary at '%s'. Used default dictionary with fruits\n", path)
		return dict.Fruits
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return []string{}
	}

	return strings.Split(string(content), "\n")
}

func main() {
	var dictPath string
	flag.StringVar(&dictPath, "dict", "", "Provide a path to some dictionary with custom words")
	flag.Parse()

	tree := trie.New()
	dictionary := getDictionary(dictPath)
	tree.InsertMany(dictionary)

	word := os.Args[len(os.Args)-1]

	results := tree.Autocomplete(word)

	fmt.Printf("Autocomplete results:\n")
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
}
