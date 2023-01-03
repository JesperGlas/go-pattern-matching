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
	data, err := os.ReadFile("data/huge.txt")
	if err != nil {
		log.Fatal(err)
	}
	sequence = string(data)
	log.Printf("Successfully loaded sequnce of %d bases..", utf8.RuneCountInString(sequence))

	fmt.Println()
	log.Println("Naive pattern matching:")
	start := time.Now()
	matches, checks := naive.CountOccurence(pattern, sequence)
	duration := time.Since(start)
	log.Printf("Status: %d/%d (matches|checks)\n", matches, checks)
	log.Printf("Runtime: %d ms\n", duration.Microseconds())

	fmt.Println()
	log.Println("Naive pattern matching (Concurrent batches):")
	start = time.Now()
	matches, checks = naive.CountOccurenceConcurrent(pattern, sequence, 10000)
	duration = time.Since(start)
	log.Printf("Status: %d/%d (matches|checks)\n", matches, checks)
	log.Printf("Runtime: %d ms\n", duration.Microseconds())
}
