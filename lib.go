package englify

import (
	"strings"
)

const (
	start      = iota
	skipTag    = iota
	skipEscape = iota
	buildWord  = iota
)

func collectWords(input string) []string {
	var word string
	words := make([]string, 0)
	state := start
	for _, char := range input {
		switch state {
		case start:
			switch char {
			case '<':
				state = skipTag
			case '&':
				state = skipEscape
			default:
				state = buildWord
				word += string(char)
			}
		case skipTag:
			switch char {
			case '>':
				state = buildWord
			default:
				state = skipTag
			}
		case skipEscape:
			switch char {
			case ';':
				state = buildWord
			default:
				state = skipEscape
			}
		case buildWord:
			switch char {
			case '<':
				state = skipTag
				words = addIfNonEmpty(words, word)
				word = ""
			case '&':
				state = skipEscape
			default:
				if isPunctuation(char) {
					state = buildWord
					words = addIfNonEmpty(words, word)
					word = ""
				} else {
					word += string(char)
				}
			}
		}
	}
	words = addIfNonEmpty(words, word)
	return words
}

func isPunctuation(char rune) bool {
	punctuations := " ,.;\r\n"
	if strings.Index(punctuations, string(char)) >= 0 {
		return true
	}
	return false
}

func addIfNonEmpty(ws []string, w string) []string {
	if w != "" {
		return append(ws, w)
	}
	return ws
}
