package go_thousands_separator

import (
	"strconv"
	"strings"
)

func ThousandsSeparator(value int) string {
	var result strings.Builder

	s := strconv.Itoa(value)
	n := len(s)

	if n < 3 {
		return s
	}

	for i := 0; i < n; i++ {
		result.WriteByte(s[i])

		if (n-i-1)%3 == 0 && i != n-1 {
			result.WriteByte('.')
		}
	}

	return result.String()
}
