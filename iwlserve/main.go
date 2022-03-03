package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	"github.com/coding-robots/goiwl/bayes"
	"github.com/coding-robots/goiwl/tokenizer"
)

var (
	serverAddr  = flag.String("s", "", "HTTP service address and port (e.g. ':8080'). If empty, use FastCGI.")
	tryClassify = flag.Bool("try", false, "read input from stdin and classify it")

	classifier *bayes.Bayes

	siteBaseURL = "https://iwl.me"
)

func classifyStdin() {
	log.SetFlags(0)
	text, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading stdin: %s", err)
	}
	t := tokenizer.WordsSpecials(string(text))

	ranks := classifier.Rank(t)
	res := make([]string, len(ranks))
	for i, v := range ranks {
		res[i] = fmt.Sprintf("%f\t%s", v, classifier.CategoryForIndex(i))
	}
	sort.Strings(res)
	for _, s := range res {
		fmt.Println(s)
	}
	/*
		t = tokenizer.WordsSpecials(string(text))
		author, rank := b.Classify(t)
		fmt.Printf("%s (%f)\n", author, rank)
	*/
}

func main() {
	flag.Parse()
	t := time.Now()
	var err error
	f, err := embeddedFS.Open("data/data.gobz")
	if err != nil {
		log.Fatal(err)
	}
	classifier, err = bayes.LoadReader(f)
	if err != nil {
		log.Fatalf("error loading data: %s", err)
	}
	f.Close()

	if *tryClassify {
		classifyStdin()
		return
	}
	fmt.Printf("Loaded in %s\n", time.Now().Sub(t))

	Serve(*serverAddr)
}
