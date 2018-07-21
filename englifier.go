package englifier

import (
	"strings"
)

const (
	start = iota
	skipTag
	skipEscape
	buildWord
)

// CollectWords is a state machine to parse HTML
// and collect words from it.
// Tags and HTML excape texts are ignored.
func CollectWords(input string) []string {
	var word string
	words := make([]string, 0)
	state := start
	actualTag := ""
	inBody := false
	for _, char := range input {
		switch state {
		case start:
			switch char {
			case '<':
				state = skipTag
			default:
				// skip
			}
		case skipTag:
			switch char {
			case '>':
				state = buildWord
				actualTag = ""
			default:
				state = skipTag
				actualTag += string(char)
				if actualTag == "body" || strings.HasPrefix(actualTag, "body ") {
					inBody = true
				}
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
				words = addIfNonEmpty(words, word, inBody)
				word = ""
			case '&':
				state = skipEscape
			default:
				if isPunctuation(char) {
					state = buildWord
					words = addIfNonEmpty(words, word, inBody)
					word = ""
				} else {
					word += string(char)
				}
			}
		}
	}
	words = addIfNonEmpty(words, word, inBody)
	return words
}

func isPunctuation(char rune) bool {
	punctuations := " ,.;:!?*#[](){}\"'\r\n"
	if strings.Index(punctuations, string(char)) >= 0 {
		return true
	}
	return false
}

func addIfNonEmpty(ws []string, w string, b bool) []string {
	if w != "" && b {
		return append(ws, strings.ToLower(w))
	}
	return ws
}
