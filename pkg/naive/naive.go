package naive

func NaiveMatching(pattern string, sequence string) (bool, int, int) {
	checks := 0
	matches := 0
	for i := 0; i < len(sequence)-len(pattern); i++ {
		checks++
		if sequence[i:i+len(pattern)] == pattern {
			matches++
		}
	}
	return matches > 0, matches, checks
}
