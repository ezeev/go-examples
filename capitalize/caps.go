package capitalize

import (
	"strings"
)

func CapWords(sent string) string {

	tokens := strings.Split(sent, " ")

	newTokens := make([]string, len(tokens))

	for i, v := range tokens {
		token := strings.Replace(v, string(v[0]), strings.ToUpper(string(v[0])), 1)
		newTokens[i] = token
	}

	return strings.Join(newTokens, " ")

}
