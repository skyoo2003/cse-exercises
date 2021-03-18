package decryptstringfromalphabettointegermapping

import (
	"strconv"
)

func freqAlphabets(s string) string {
	var buf []rune
	i := len(s) - 1
	for i >= 0 {
		var v int
		if s[i] == '#' {
			v, _ = strconv.Atoi(s[i-2 : i])
			i -= 3
		} else {
			v, _ = strconv.Atoi(string(s[i]))
			i--
		}
		r := rune('a' + v - 1)
		buf = append([]rune{r}, buf...)
	}
	return string(buf)
}
