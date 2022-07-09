package reverseInteger

import (
	errors "github.com/pkg/errors"
	"math"
)

var ErrHelp = errors.New("Provided help")
var ErrOverflow = errors.New("integer overflow")

func ReverseInteger() error {

	return nil
}

func reverseInteger(x int) int {

	var err error
	var t []int
	res := 0
	for x != 0 {
		t = append(t, x%10)
		x /= 10
	}

	for i, elem := range t {
		res, err = add32(res, elem*int(math.Pow10(len(t)-i-1)))

		if err != nil {
			return 0
		}
	}

	return res
}

func add32(left, right int) (int, error) {
	if right > 0 {
		if left > math.MaxInt32-right {
			return 0, ErrOverflow
		}
	} else {
		if left < math.MinInt32-right {
			return 0, ErrOverflow
		}
	}
	return left + right, nil
}
