package bayes

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/coding-robots/goiwl/tokenizer"
)

func (b *Bayes) trainToken(catIndex int, word string) {
	// Increment totals.
	b.totals[catIndex]++
	// Increment tokens.
	t, ok := b.tokens[word]
	if !ok {
		t = make([]uint32, catIndex+1)
		b.tokens[word] = t
	} else if len(t) <= catIndex {
		tmp := make([]uint32, catIndex+1)
		copy(tmp, t)
		t = tmp
		b.tokens[word] = t
	}
	t[catIndex]++
}

func (b *Bayes) RemoveExtremes() {
	cutoff := len(b.categories)-len(b.categories)/5
	for i, token := range b.tokens {
		// Count zeros.
		zeros := 0
		for _, n := range token {
			if n == 0 {
				zeros++
			}
		}
		if zeros >= cutoff {
			delete(b.tokens, i)
			// ...but leave totals as is.
		}
	}
}

func (b *Bayes) Train(category string, tz tokenizer.Tokenizer) {
	catIndex := b.categoryIndex(category)
	for tz.Next() {
		b.trainToken(catIndex, tz.Token())
	}
}

func (b *Bayes) TrainFile(category string, filename string) error {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	b.Train(category, tokenizer.WordsSpecials(string(text)))
	return nil
}

func (b *Bayes) trainAuthorDir(dir string) error {
	author := filepath.Base(dir)
	files, err := filepath.Glob(filepath.Join(dir, "*.txt"))
	if err != nil {
		return err
	}
	if files == nil {
		return fmt.Errorf("No *.txt files found in %s", dir)
	}
	for _, f := range files {
		log.Printf("Training as %s file %s", author, f)
		if err := b.TrainFile(author, f); err != nil {
			return err
		}
	}
	return nil
}

// TrainDirs trains the classifier on the directory contents.
// The directory must contain one directory per author, which
// must contain texts to train on in "*.txt" files.
//
// Example layout:
//
//	train-data\
//		Conan Doyle\
//			holmes.txt
//			lost-world.txt
//		Christie\
//			poirot.txt
//			marple.txt
//
func (b *Bayes) TrainDirs(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	for {
		fi, err := d.Readdir(1)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if !fi[0].IsDir() {
			continue
		}
		if err := b.trainAuthorDir(filepath.Join(dir, fi[0].Name())); err != nil {
			return err
		}
	}
}
