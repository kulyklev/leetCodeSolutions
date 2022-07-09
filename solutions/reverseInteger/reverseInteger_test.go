package reverseInteger

import "testing"

func TestReverseInteger(t *testing.T) {
	/* Case 1 */
	got := reverseInteger(123)
	want := 321

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	/* Case 2 */
	got = reverseInteger(120)
	want = 21

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	///* Case 3 */
	got = reverseInteger(102)
	want = 201

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	///* Case 4 */
	got = reverseInteger(-123)
	want = -321

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	///* Case 5 */
	got = reverseInteger(-120)
	want = -21

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	///* Case 6 */
	got = reverseInteger(-2147483648)
	want = 0

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	///* Case 7 */
	got = reverseInteger(2147483647)
	want = 0

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	///* Case 8 */
	got = reverseInteger(0)
	want = 0

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
