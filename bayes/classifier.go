package bayes

import (
	"fmt"
	"math"
	"sort"

	"github.com/coding-robots/goiwl/tokenizer"
)

const Debugging = false

const noWordRating = 0.4

// Rank returns a slice of probabilities indexed by category.
func (b *Bayes) Rank(tz tokenizer.Tokenizer) []float64 {
	// Initialize ratings map.
	ratings := make([][]float64, len(b.categories))
	var words []string
	if Debugging {
		words = make([]string, 0, 10)
	}
	for i := range b.categories {
		ratings[i] = make([]float64, 0, 10)
	}
	sumTotals := b.sumTotals()
	for tz.Next() {
		tokenCounts, ok := b.tokens[tz.Token()]
		if Debugging {
			words = append(words, tz.Token())
		}
		if !ok {
			// Never seen this word.
			for i := range b.categories {
				ratings[i] = append(ratings[i], noWordRating)
			}
			continue
		}
		// Calculate tokens sum.
		var sumToken float64
		for _, v := range tokenCounts {
			sumToken += float64(v)
		}

		for catIndex, catTotal := range b.totals {
			if catIndex >= len(tokenCounts) {
				ratings[catIndex] = append(ratings[catIndex], 0.2)
				continue
			}
			cnt := tokenCounts[catIndex]
			if cnt == 0 {
				ratings[catIndex] = append(ratings[catIndex], 0.2)
				continue
			}
			thisProb := float64(cnt) / float64(catTotal)
			otherProb := (sumToken - float64(cnt)) / (sumTotals - float64(catTotal))
			//rating := math.Log(thisProb) / math.Log(thisProb + otherProb)
			rating := thisProb / (thisProb + otherProb)
			if rating < 0.20 {
				rating = 0.20
			} else if rating > 0.70 {
				rating = 0.70
			}
			ratings[catIndex] = append(ratings[catIndex], rating)
		}
	}

	// Calculate ranks.
	ranks := make([]float64, len(ratings))
	for catIndex, probs := range ratings {
		if Debugging {
			fmt.Printf("%s: \n", b.CategoryForIndex(catIndex))
			for i, v := range probs {
				fmt.Printf("\t%f\t%s\n", v, words[i])
			}
		}
		// Get rid of 1 highest and 1 lowest rating,
		// then leave only 10 lowest and 10 highest results.
		if len(probs) > 6 {
			sort.Float64s(probs)
			probs = probs[2:len(probs)-2]
			if len(probs) > 80 {
				probs = append(probs[:40], probs[len(probs)-40:]...)
			}
		}

		nth := 1.0 / float64(len(probs))
		p := 1.0
		q := 1.0
		for _, v := range probs {
			p *= 1.0 - v
			q *= v
		}
		p = 1.0 - math.Pow(p, nth)
		q = 1.0 - math.Pow(q, nth)
		s := (p - q) / (p + q)
		ranks[catIndex] = (1 + s) / 2
	}
	return ranks
}

func (b *Bayes) Classify(tz tokenizer.Tokenizer) (category string, score float64) {
	ranks := b.Rank(tz)
	index := 0
	// Find highest rank.
	for i, r := range ranks {
		if r > score {
			score = r
			index = i
		}
	}
	category = b.CategoryForIndex(index)
	return
}
