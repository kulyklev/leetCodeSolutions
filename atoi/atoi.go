package atoi

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	tmpStr := strings.TrimSpace(s)
	var sb strings.Builder

	for i, rn := range tmpStr {

		if i == 0 && (rn == '+' || rn == '-') {
			sb.WriteRune(rn)
			continue
		}

		if unicode.IsDigit(rn) {
			sb.WriteRune(rn)
		} else {
			break
		}

	}

	res, err := strconv.Atoi(sb.String())

	if err != nil {
		fmt.Println("Convert string to int")
		fmt.Println(err)
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}
