package atoi

import "testing"

func TestAtoi(t *testing.T) {
	// Case 1
	got := myAtoi("123")
	want := 123

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	// Case 2
	got = myAtoi("-123")
	want = -123

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	// Case 3
	got = myAtoi("    123")
	want = 123

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	// Case 4
	got = myAtoi("   -123")
	want = -123

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	// Case 5
	got = myAtoi("0032")
	want = 32

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	// Case 6
	got = myAtoi("qwerty")
	want = 0

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	// Case 7
	got = myAtoi("1234 qwerty")
	want = 1234

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	// Case 8
	got = myAtoi("-1234 qwerty")
	want = -1234

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	// Case 9
	got = myAtoi("-91283472332")
	want = -2147483648

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	// Case 10
	got = myAtoi("91283472332")
	want = 2147483647

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
