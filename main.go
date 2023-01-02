package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"unicode/utf8"

	"github.com/JesperGlas/go-pattern-matching/pkg/naive"
)

func main() {
	fmt.Println("Running pattern matching algorithms..")

	// Set sought after pattern
	var pattern string = "atg"
	log.Printf("Pattern: %s\n", pattern)

	// Load sequence as string from file
	var sequence string
	data, err := os.ReadFile("data/large.txt")
	if err != nil {
		log.Fatal(err)
	}
	sequence = string(data)
	log.Printf("Successfully loaded sequnce of %d bases..", utf8.RuneCountInString(sequence))

	fmt.Println()
	log.Println("Naive pattern matching:")
	start := time.Now()
	status, matches, checks := naive.NaiveMatching(pattern, sequence)
	duration := time.Since(start)
	log.Printf("Status: %v (%d/%d matches|checks)\n", status, matches, checks)
	log.Printf("Runtime: %d ms\n", duration.Microseconds())
}
