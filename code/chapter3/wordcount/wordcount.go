package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yuezhilanyi/goinaction/code/chapter3/words"
)

// main is the entry point for the application.
func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
