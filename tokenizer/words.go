package tokenizer

import (
	"strings"
	"unicode"
)

type words struct {
	input string
	word  string
}

// Words returns a Tokenizer that tokenizes input into words.
func Words(input string) Tokenizer {
	return &words{input: input}
}

func isWordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '%' || r == '$'
}

func isNotWordRune(r rune) bool {
	return !isWordRune(r)
}

func (w *words) Next() bool {
	s := w.input
	// Skip non-word runes.
	start := strings.IndexFunc(s, isWordRune)
	if start == -1 {
		return false
	}
	s = s[start:]
	// Get index of next non-word rune.
	end := strings.IndexFunc(s, isNotWordRune)
	if end == -1 {
		end = len(s)
	}
	if end == 0 {
		return false
	}
	w.word = s[0:end]
	w.input = s[end:]
	return true
}

func (w *words) Token() string {
	return strings.ToLower(w.word)
}
