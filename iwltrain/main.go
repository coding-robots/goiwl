package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coding-robots/goiwl/bayes"
)

var (
	inDir   = flag.String("d", "train-data", "input directory")
	outFile = flag.String("o", "data.gobz", "output data file")
)

func main() {
	log.SetFlags(0)
	flag.Parse()
	if *inDir == "" || *outFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	t := time.Now()
	b := bayes.New()
	if err := b.TrainDirs(*inDir); err != nil {
		log.Fatalf("error training: %s", err)
	}
	b.RemoveExtremes()
	if err := b.WriteFile(*outFile); err != nil {
		log.Fatalf("error writing trained data: %s", err)
	}
	fmt.Printf("Done in %s\n", time.Now().Sub(t))
}
