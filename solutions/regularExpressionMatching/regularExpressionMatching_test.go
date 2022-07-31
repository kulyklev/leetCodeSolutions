package regularExpressionMatching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestIsMatch(t *testing.T) {

	t.Log("Given the need to check regexp match.")
	{
		for scenario, fn := range map[string]func(t *testing.T){
			"match multiple single characters": testMatchMultipleCharacters,
		} {
			t.Run(scenario, func(t *testing.T) {
				fn(t)
			})
		}
	}
}

func testMatchMultipleCharacters(t *testing.T) {

	type testCases struct {
		description  string
		inputString  string
		inputPattern string
		expected     bool
	}

	testID := 0
	t.Logf("\t\tTest %d:\tWhen matching pattern with different patterns with placeholders.", testID)
	{
		for _, scenario := range []testCases{
			{
				description:  "No pattern match",
				inputString:  "qwe",
				inputPattern: "qwe",
				expected:     true,
			},
			{
				description:  "No pattern mismatch.",
				inputString:  "qwe",
				inputPattern: "qwerty",
				expected:     false,
			},
			{
				description:  "Match single character",
				inputString:  "a",
				inputPattern: ".",
				expected:     true,
			},
			{
				description:  "Two duplicate chars",
				inputString:  "aa",
				inputPattern: "..",
				expected:     true,
			},
			{
				description:  "Two different chars",
				inputString:  "ab",
				inputPattern: "..",
				expected:     true,
			},
			{
				description:  "Placeholders at edges",
				inputString:  "cabd",
				inputPattern: ".ab.",
				expected:     true,
			},
			{
				description:  "Base zero or many place holders",
				inputString:  "aa",
				inputPattern: "a*",
				expected:     true,
			},
			{
				description:  "Zero or many place holders 1",
				inputString:  "aaaaa",
				inputPattern: "a*",
				expected:     true,
			},
			{
				description:  "Zero or many place holders 2",
				inputString:  "asdfbvxcgyurtfghj",
				inputPattern: ".*",
				expected:     true,
			},
			{
				description:  "Zero or many place holders 3",
				inputString:  "tab",
				inputPattern: ".abc*",
				expected:     true,
			},
			{
				description:  "Zero or many place holders 4",
				inputString:  "aaa",
				inputPattern: "aa",
				expected:     false,
			},
			{
				description:  "Zero or many place holders 5",
				inputString:  "abcd",
				inputPattern: "d*",
				expected:     false,
			},
			{
				description:  "Zero or many place holders 6",
				inputString:  "aaa",
				inputPattern: "aaaa",
				expected:     false,
			},
			{
				description:  "Zero or many place holders 7",
				inputString:  "aaa",
				inputPattern: "ab*a",
				expected:     false,
			},
			{
				description:  "Zero or many place holders 8",
				inputString:  "aaa",
				inputPattern: "ab*ac*a",
				expected:     true,
			},
			{
				description:  "Zero or many place holders 9",
				inputString:  "aaba",
				inputPattern: "ab*a*c*a",
				expected:     false,
			},
			{
				description:  "Duplicate chars without placeholders",
				inputString:  "aa",
				inputPattern: "a",
				expected:     false,
			},
			{
				description:  "Mississippi case",
				inputString:  "mississippi",
				inputPattern: "mis*is*p*.",
				expected:     false,
			},
			{
				description:  "Mississippi case 2",
				inputString:  "mississippi",
				inputPattern: "mis*is*ip*.",
				expected:     true,
			},
		} {
			t.Run(scenario.description, func(t *testing.T) {
				actual := isMatch(scenario.inputString, scenario.inputPattern)

				assert.Equal(t, scenario.expected, actual)
			})
		}
	}
}
