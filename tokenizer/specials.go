package tokenizer

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	wordCountTokenFmt      = "WC=%d"
	commaCountTokenFmt     = "CC=%d"
	semicolonCountTokenFmt = "SCC=%d"
	hasQuoteToken          = "HQ"
	hasDashToken           = "HD"
	sentenceShapeToken     = "SH=%s"
)

type specials struct {
	sentences  Tokenizer
	tokens     []string
	tokenIndex int
}

// Specials returns a Tokenizer that generates special tokens,
// such as word count, comma count, etc., for each sentence of
// input.
func Specials(input string) Tokenizer {
	s := new(specials)
	s.sentences = Sentences(input)
	s.tokens = make([]string, 0, 5)
	s.tokenIndex = -1
	return s
}

func wordCount(s string) int {
	n := 0
	w := Words(s)
	for w.Next() {
		n++
	}
	return n
}

func commaCount(s string) int {
	return strings.Count(s, ",")
}

func semicolonCount(s string) int {
	return strings.Count(s, ";")
}

func hasQuote(s string) bool {
	return strings.IndexRune(s, '"') != -1
}

func hasDash(s string) bool {
	return strings.Index(s, "- ") != -1 || strings.Index(s, "--") != -1
}

func sentenceShape(s string) string {
	shape := make([]byte, 0, 10)
	prevType := byte(0)
	nw := 0
	for _, r := range s {
		if isWordRune(r) {
			if prevType != 'W' {
				if prevType != 0 {
					shape = append(shape, prevType)
				}
				nw++
			}
			prevType = 'W'
			continue
		}
		newType := byte(0)
		switch r {
		case ',', '.', '?', ':', ';', '!':
			newType = byte(r)
		}
		if nw > 0 {
			if newType != 0 {
				shape = append(shape, strconv.FormatInt(int64(nw), 36)...)
				nw = 0
			}
		}
		if prevType != 0 && prevType != 'W' {
			shape = append(shape, prevType)
		}
		prevType = newType
	}
	if nw > 0 {
		shape = append(shape, strconv.FormatInt(int64(nw), 36)...)
	}
	if prevType != 0 && prevType != 'W' {
		shape = append(shape, prevType)
	}
	return string(shape)
}

func (s *specials) refill() bool {
	s.tokens = s.tokens[:0]
	s.tokenIndex = 0
	if !s.sentences.Next() {
		return false
	}
	sentence := s.sentences.Token()
	// Word count.
	if n := wordCount(sentence); n > 3 {
		s.tokens = append(s.tokens, fmt.Sprintf(wordCountTokenFmt, n))
	}
	// Comma count.
	if n := commaCount(sentence); n > 0 {
		s.tokens = append(s.tokens, fmt.Sprintf(commaCountTokenFmt, n))
	}
	// Semicolon count.
	if n := semicolonCount(sentence); n > 0 {
		s.tokens = append(s.tokens, fmt.Sprintf(semicolonCountTokenFmt, n))
	}
	// Has quote?
	if hasQuote(sentence) {
		s.tokens = append(s.tokens, hasQuoteToken)
	}
	// Has dash?
	if hasDash(sentence) {
		s.tokens = append(s.tokens, hasDashToken)
	}
	// Sentence shape.
	if shape := sentenceShape(sentence); shape != "" {
		s.tokens = append(s.tokens, fmt.Sprintf(sentenceShapeToken, shape))
	}
	// To add more weight to specials, add them two times.
	if len(s.tokens) > 0 {
		s.tokens = append(s.tokens, s.tokens...)
	}
	return true
}

func (s *specials) Next() bool {
	s.tokenIndex++
	for s.tokenIndex >= len(s.tokens) {
		if !s.refill() {
			return false
		}
	}
	return true
}

func (s *specials) Token() string {
	return s.tokens[s.tokenIndex]
}
