package reversewords

import "strings"

func ReverseWords(str string) string {
	result := make([]string, 0, len(str)/2)
	end := len(str) - 1
	for i := len(str) - 1; i >= 0; i-- {
		next := i - 1
		if str[i] != ' ' {
			if i+1 < len(str) && str[i+1] == ' ' {
				end = i
			}
			if next < 0 || str[next] == ' ' {
				result = append(result, str[i:end+1])
				end = next
			}
		}
	}
	return strings.Join(result, " ")
}
