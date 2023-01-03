package naive

import (
	"sync"
)

/**
* Desc: Function that counts occurences of a pattern in a longer sequence
*
* Args:
*	pattern (string)	- sought after pattern
*	sequence (string)	- sequence of characters
*
* Return:
*	(int) - number of matches
*	(int) - number of checks
 */
func CountOccurence(pattern string, sequence string) (int, int) {
	checks := 0
	matches := 0
	for i := 0; i < len(sequence)-len(pattern); i++ {
		checks++
		if sequence[i:i+len(pattern)] == pattern {
			matches++
		}
	}
	return matches, checks
}

/**
* Struct for a mutex locked counter
 */
type mutexCounter struct {
	lock  sync.Mutex
	value int
}

/**
* Desc: Initialize function for the mutexCounter struct
*
* Return:
*	(*mutexCounter) - pointer to a new mutexCounter with value 0
 */
func initCounter() *mutexCounter {
	return &mutexCounter{value: 0}
}

/**
* Desc: Function to increment an instance of the mutexCounter struct
*
* Args:
*	counter (*mutexCounter)	- pointer to a mutexCounter instance
*	value (int)				- value to increment with
 */
func incrementCounter(counter *mutexCounter, value int) {
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.value += value
}

/*
*
  - Desc: Function to count the occurences of a pattern in a longer sequence
    The function does not count edge cases between batches (yet)

*
* Args:
*	pattern (string)	- sought after pattern
*	sequence (string)	- sequence of characters
*	batchSize (int)		- the size of the sequence each worker will check
*
* Return:
*	(int) - total number of occurences
*	(int) - total number of checks performed
*/
func CountOccurenceConcurrent(pattern string, sequence string, batchSize int) (int, int) {
	// init counters for checks and matches
	sharedChecks := initCounter()
	sharedMatches := initCounter()

	// define a waitgroup for goroutines
	var wg sync.WaitGroup

	// create batches for goroutines
	for i := 0; i < len(sequence); i += batchSize {
		// define start index
		start := i

		end := i + batchSize + len(pattern)
		if end > len(sequence) {
			end = len(sequence)
		}

		// add worker to waitgroup and start worker
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			matches, checks := CountOccurence(pattern, sequence[start:end])
			incrementCounter(sharedChecks, checks)
			incrementCounter(sharedMatches, matches)
		}(&wg)
	}

	// wait for all goroutines to finish before return
	wg.Wait()
	return sharedMatches.value, sharedChecks.value
}
