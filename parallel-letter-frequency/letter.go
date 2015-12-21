package letter

/*ConcurrentFrequency counts word occurances using parallel processes*/
func ConcurrentFrequency(wordList []string) FreqMap {
	// Map the frequency function over all the words
	channel := make(chan FreqMap, len(wordList))
	for _, words := range wordList {
		go func(w string) { channel <- Frequency(w) }(words)
	}

	// Reduce the results to a single map
	frequency := FreqMap{}
	for range wordList {
		for key, value := range <-channel {
			frequency[key] += value
		}
	}
	return frequency
}
