package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/afabian80/englifier"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Cannot read input")
	}
	words := englifier.CollectWords(string(bytes))
	fmt.Println(words)
}
