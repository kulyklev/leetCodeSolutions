package regularExpressionMatching

import "testing"

func TestIsMatch(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"no pattern match":    testNoPatternMatch,
		"no pattern mismatch": testNoPatternMismatch,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testNoPatternMatch(t *testing.T) {
	got := isMatch("qwe", "w")
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func testNoPatternMismatch(t *testing.T) {
	got := isMatch("qwe", "w")
	want := false

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
