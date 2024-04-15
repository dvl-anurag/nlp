package nlp

import (
	"regexp"
	"strings"
)

func tokenise(textStr string) []string {
	reg := regexp.MustCompile(`[\s,.!?]+`)

	tokens := reg.Split(textStr, -1)

	var result []string
	for _, token := range tokens {
		if strings.TrimSpace(token) != "" {
			result = append(result, token)
		}
	}

	return result
}
