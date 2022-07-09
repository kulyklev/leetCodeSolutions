package palindromeNumber

import "testing"

func TestIsPalindromeNumber(t *testing.T) {
	// Case 1
	got := isPalindrome(121)
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}

	// Case 2
	got = isPalindrome(-121)
	want = false

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}

	// Case 3
	got = isPalindrome(1234)
	want = false

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
