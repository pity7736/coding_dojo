package mergealternately

import "strings"

func MergeAlternately(word1, word2 string) string {
	b := strings.Builder{}
	rune1 := []rune(word1)
	rune1Len := len(rune1)
	rune2 := []rune(word2)
	rune2Len := len(rune2)
	maxLen := max(rune1Len, rune2Len)
	b.Grow(rune1Len + rune2Len)
	for i := range max(maxLen) {
		if i < rune1Len {
			b.WriteRune(rune1[i])
		}
		if i < rune2Len {
			b.WriteRune(rune2[i])
		}
	}
	return b.String()
}
