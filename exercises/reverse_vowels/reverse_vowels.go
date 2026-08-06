package reversevowels

import "strings"

func ReverseVowels(str string) string {
	left := 0
	strRunes := []rune(str)
	right := len(strRunes) - 1
	for left <= right {
		if isVowel(strRunes[left]) && isVowel(strRunes[right]) {
			strRunes[left], strRunes[right] = strRunes[right], strRunes[left]
			left++
			right--
		} else {
			if !isVowel(strRunes[left]) {
				left++
			}
			if !isVowel(strRunes[right]) {
				right--
			}
		}
	}
	return string(strRunes)
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aoeuiAOEUIáéíóú", r)

}
