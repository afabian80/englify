package main

import (
	"fmt"

	"github.com/afabian80/englify"
)

func main() {
	s := `<html><body class="apple">the</body></html>`
	words := englify.CollectWords(s)
	fmt.Println(words)
}
