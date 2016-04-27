package bayes

import (
	"compress/gzip"
	"encoding/gob"
	"os"
)

type Bayes struct {
	categories []string
	// {category_index: 0, ...}
	totals map[int]float64
	// {"word":{0:33, ...}}
	tokens map[string][]uint32
}

type serializedBayes struct {
	Categories []string
	Totals     map[int]float64
	Tokens     map[string][]uint32
}

func New() *Bayes {
	b := new(Bayes)
	b.categories = make([]string, 0, 52)
	b.totals = make(map[int]float64)
	b.tokens = make(map[string][]uint32)
	return b
}

func LoadFile(filename string) (*Bayes, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	eb := new(serializedBayes)
	z, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	if err := gob.NewDecoder(z).Decode(eb); err != nil {
		return nil, err
	}

	return &Bayes{eb.Categories, eb.Totals, eb.Tokens}, nil
}

func (b *Bayes) WriteFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	eb := &serializedBayes{
		b.categories,
		b.totals,
		b.tokens,
	}
	z, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		f.Close()
		return err
	}
	if err := gob.NewEncoder(z).Encode(eb); err != nil {
		z.Close()
		f.Close()
		return err
	}
	if err := z.Close(); err != nil {
		f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func (b *Bayes) categoryIndex(cat string) int {
	// Try to find category.
	for i, v := range b.categories {
		if v == cat {
			return i
		}
	}
	// Not found, append.
	b.categories = append(b.categories, cat)
	return len(b.categories) - 1
}

func (b *Bayes) CategoryForIndex(catIndex int) string {
	return b.categories[catIndex]
}

func (b *Bayes) Categories() []string {
	return b.categories
}

func (b *Bayes) NumberOfTokens() int {
	return len(b.tokens)
}

func (b *Bayes) sumTotals() (sum float64) {
	//TODO store this thing instead of calculating it.
	for _, v := range b.totals {
		sum += v
	}
	return
}
