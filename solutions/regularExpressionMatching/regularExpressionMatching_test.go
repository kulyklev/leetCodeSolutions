package regularExpressionMatching

import "testing"

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestIsMatch(t *testing.T) {

	t.Log("Given the need to check regexp match.")
	{
		for scenario, fn := range map[string]func(t *testing.T){
			"no pattern match":                 testNoPatternMatch,
			"no pattern mismatch":              testNoPatternMismatch,
			"match one single character":       testMatchSingleCharacter,
			"match multiple single characters": testMatchMultipleCharacters,
		} {
			t.Run(scenario, func(t *testing.T) {
				fn(t)
			})
		}
	}
}

func testNoPatternMatch(t *testing.T) {

	testID := 0
	t.Logf("\t\tTest %d:\tWhen matching without pattern.", testID)
	{
		inputString := "qwe"
		pattern := "qwe"
		got := isMatch(inputString, pattern)
		want := true

		if got != want {
			t.Fatalf("\t%s\tTest %d:\tExpected: %t. String: %s. Does not match: %s", failed, testID, want, inputString, pattern)
		}

		t.Logf("\t%s\tTest %d:\tString matches pattern.", success, testID)
	}
}

// This case is tricky
func testNoPatternMismatch(t *testing.T) {

	testID := 0
	t.Logf("\t\tTest %d:\tWhen not matching without pattern.", testID)
	{
		inputString := "qwe"
		pattern := "qwerty"
		got := isMatch(inputString, pattern)
		want := false

		if got != want {
			t.Fatalf("\t%s\tTest %d:\tExpected: %t. String: %s. Match pattern: %s", failed, testID, want, inputString, pattern)
		}

		t.Logf("\t%s\tTest %d:\tString mismatches.", success, testID)
	}
}

func testMatchSingleCharacter(t *testing.T) {
	testID := 0
	t.Logf("\t\tTest %d:\tWhen matching pattern with singe character.", testID)
	{
		inputString := "a"
		pattern := "."
		got := isMatch(inputString, pattern)
		want := true

		if got != want {
			t.Fatalf("\t%s\tTest %d:\tExpected: %t. String: \"%s\" should match pattern: \"%s\".", failed, testID, want, inputString, pattern)
		}

		t.Logf("\t%s\tTest %d:\tPattern matches.", success, testID)
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
		} {
			t.Run(scenario.description, func(t *testing.T) {
				actual := isMatch(scenario.inputString, scenario.inputPattern)

				if actual != scenario.expected {

					if scenario.expected {
						t.Fatalf("\t%s\tTest %d:\tExpected: %t. String: \"%s\" should match pattern: \"%s\".", failed, testID, scenario.expected, scenario.inputString, scenario.inputPattern)
					} else {
						t.Fatalf("\t%s\tTest %d:\tExpected: %t. String: \"%s\" should NOT match pattern: \"%s\".", failed, testID, scenario.expected, scenario.inputString, scenario.inputPattern)
					}
				}

				t.Logf("\t%s\tTest %d:\tPattern matches.", success, testID)
				/*require.Error(t, err)
				assert.Equal(t, scenario.expectedError, err.Error())*/
			})
		}
	}
}
