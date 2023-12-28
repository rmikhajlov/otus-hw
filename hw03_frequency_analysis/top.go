package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	if len(text) == 0 {
		return make([]string, 0)
	}
	stringsInText := strings.Fields(text)

	stringCount := make(map[string]int)

	for i := range stringsInText {
		stringCount[stringsInText[i]]++
	}

	countLen := len(stringCount)
	result := make([]string, 0, countLen)

	for key := range stringCount {
		result = append(result, key)
	}

	sort.Slice(result, func(i, j int) bool {
		if stringCount[result[i]] == stringCount[result[j]] {
			return result[i] < result[j]
		}
		return stringCount[result[i]] > stringCount[result[j]]
	})

	if countLen < 10 {
		return result[:countLen]
	}

	return result[:10]
}
