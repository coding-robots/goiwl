package bayes

import (
	"compress/gzip"
	"encoding/gob"
	"io"
	"os"
)

type Bayes struct {
	Categories []string
	// {category_index: 0, ...}
	Totals map[int]float64
	// {"word":{0:33, ...}}
	Tokens map[string][]uint32

	sumTotals float64
}

func New() *Bayes {
	return &Bayes{
		Categories: make([]string, 0, 52),
		Totals:     make(map[int]float64),
		Tokens:     make(map[string][]uint32),
	}
}

func LoadFile(filename string) (*Bayes, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LoadReader(f)
}

func LoadReader(r io.Reader) (*Bayes, error) {
	z, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	var b Bayes
	if err := gob.NewDecoder(z).Decode(&b); err != nil {
		return nil, err
	}
	b.calcSumTotals()
	return &b, nil
}

func (b *Bayes) WriteFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	z, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		f.Close()
		return err
	}
	if err := gob.NewEncoder(z).Encode(b); err != nil {
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
	for i, v := range b.Categories {
		if v == cat {
			return i
		}
	}
	// Not found, append.
	b.Categories = append(b.Categories, cat)
	return len(b.Categories) - 1
}

func (b *Bayes) CategoryForIndex(catIndex int) string {
	return b.Categories[catIndex]
}

func (b *Bayes) NumberOfTokens() int {
	return len(b.Tokens)
}

func (b *Bayes) calcSumTotals() {
	b.sumTotals = 0
	for _, v := range b.Totals {
		b.sumTotals += v
	}
}
