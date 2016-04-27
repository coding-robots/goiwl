package tokenizer

// Tokenizer is an interface for tokenizers.
type Tokenizer interface {
	// Next processes the next token and returns true
	// if there are more tokens to process.
	Next() bool
	// Token returns the current token.
	// It must be called after Next.
	Token() string
}

type chain struct {
	tokenizers []Tokenizer
	tokIndex   int
	token      string
}

// Chain returns a Tokenizer that combines the given tokenizers into a chain.
//
// The Next method of chain gets tokens serially from the chained tokenizers
// until the last one returns false. For example, the following code:
//
// 	ch := tokenizer.Chain(tokenizer.Words("hello world"), tokenizer.Words("this is me"))
// 	for ch.Next() {
//		fmt.Printf("%s, ", ch.Token())
//	}
//
// will print:
//
//	hello, world, this, is, me, 
//
func Chain(tokenizers ...Tokenizer) Tokenizer {
	return &chain{tokenizers: tokenizers}
}

func (ch *chain) Next() bool {
	for {
		if ch.tokenizers[ch.tokIndex].Next() {
			return true
		}
		ch.tokIndex++
		if ch.tokIndex >= len(ch.tokenizers) {
			return false
		}
	}
}

func (ch *chain) Token() string {
	return ch.tokenizers[ch.tokIndex].Token()
}

// WordsSpecials returns a Tokenizer chain of Words and Specials.
func WordsSpecials(input string) Tokenizer {
	return Chain(Words(input), Specials(input))
}
