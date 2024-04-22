package nlp

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/BurntSushi/toml"
)

func TestTokeniseTable(t *testing.T) {
	tests := []struct {
		Name   string
		Text   string
		Tokens []string
	}{
		{
			Name:   "Basic test",
			Text:   "Whos many",
			Tokens: []string{"Whos", "many"},
		},
		{
			Name:   "Multiple spaces and punctuation",
			Text:   "  Hello,    World!  ",
			Tokens: []string{"Hello", "World"},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			result := Tokenize(tc.Text)
			if !reflect.DeepEqual(result, tc.Tokens) {
				t.Errorf("Test case failed for text '%s': expected %v, got %v", tc.Text, tc.Tokens, result)
			}
		})
	}
}

func readTestCasesFromTOMLFile(filename string) (map[string]TestCase, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cases []TestCase
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tc TestCase
		if err := toml.Unmarshal([]byte(scanner.Text()), &tc); err != nil {
			return nil, err
		}
		cases = append(cases, tc)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	casesMap := make(map[string]TestCase)
	for _, tc := range cases {
		casesMap[tc.Text] = tc
	}

	fmt.Println(casesMap)
	return casesMap, nil
}

type TestCase struct {
	Text   string   `toml:"text"`
	Tokens []string `toml:"tokens"`
}

func TestTokenizeFromTOML(t *testing.T) {
	file, err := os.Open("tokenize_cases.toml")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()

	d := toml.NewDecoder(file)

	var cases TCases
	_, err = d.Decode(&cases)
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, c := range cases.Cases {
		t.Run(c.Text, func(t *testing.T) {
			result := Tokenize(c.Text)

			if !reflect.DeepEqual(result, c.Tokens) {
				t.Errorf("Test case failed for text '%s': expected %v, got %v", c.Text, c.Tokens, result)
			}
		})
	}
}

type Case struct {
	Text   string   `toml:"text"`
	Tokens []string `toml:"tokens"`
}

type TCases struct {
	Cases []Case `toml:"cases"`
}
