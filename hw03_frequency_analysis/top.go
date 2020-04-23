package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

type WordID int // для простоты понимания того, что мы храним в индексе

type Word struct {
	count int
	value string
}

func Top10(s string) []string {
	index := map[string]WordID{}
	words := make([]Word, 0)
	re := regexp.MustCompile(`(?mi)([\p{L}\d]+-?[\p{L}\d]*)`)
	for _, match := range re.FindAllString(s, -1) {
		matchLow := strings.ToLower(match)
		if _, ok := index[matchLow]; !ok {
			w := Word{1, matchLow}
			index[w.value] = WordID(len(words))
			words = append(words, w)
		} else {
			id := index[matchLow]
			words[id].count++
		}
	}

	sort.Slice(words, func(i int, j int) bool {
		return words[i].count > words[j].count
	})

	// fmt.Printf("words: %v\n", words)

	top := make([]string, 0)
	for i := 0; i < 10; i++ {
		if len(words) > i {
			top = append(top, words[i].value)
		}
	}
	return top
}
