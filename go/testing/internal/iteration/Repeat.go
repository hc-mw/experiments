package iteration

import "strings"

func Repeat(ch string, n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(ch)
	}
	return sb.String()
}

func RepeatSlow(ch string, n int) string {
	var res string
	for i := 0; i < n; i++ {
		res += ch
	}
	return res
}
