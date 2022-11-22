package reverseworldsinstring3

import (
	"strings"
)

func reverseWords(s string) string {
	sl := strings.Split(s, " ")

	reversed := make([]string, 0, len(sl))

	for _, s1 := range sl {
		s1Bytes := []byte(s1)
		for b, e := 0, len(s1)-1; b < e; {
			s1Bytes[b], s1Bytes[e] = s1Bytes[e], s1Bytes[b]
			b++
			e--
		}

		reversed = append(reversed, string(s1Bytes))
	}

	return strings.Join(reversed, " ")
}
