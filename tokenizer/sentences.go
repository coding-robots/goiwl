package tokenizer

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type sentences struct {
	input    string
	sentence string
}

// Sentences returns a Tokenizer that tokenizes input into sentences.
func Sentences(input string) Tokenizer {
	return &sentences{input: input}
}

func isSentenceStartRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '"' || r == '-'
}

func isSentenceEndRune(r rune) bool {
	switch r {
	case '.', '!', '?':
		return true
	}
	return false
}

func (t *sentences) Next() bool {
	s := t.input
	// Find sentence start.
	start := strings.IndexFunc(s, isSentenceStartRune)
	if start == -1 {
		return false
	}
	s = s[start:]
	// Sentence ends with a sentence end rune.
	end := strings.IndexFunc(s, isSentenceEndRune)
	if end == -1 {
		end = len(s)
	}
	if end == 0 {
		return false
	}
	// Include as many ending marks as possible (e.g. !!!, ?!)
	for end < len(s) {
		rune, size := utf8.DecodeRuneInString(s[end:])
		if rune == utf8.RuneError {
			return false //XXX error reporting?
		}
		if !isSentenceEndRune(rune) {
			break
		}
		end += size
	}
	t.sentence = s[0:end]
	t.input = s[end:]
	return true
}

func (t *sentences) Token() string {
	return t.sentence
}
