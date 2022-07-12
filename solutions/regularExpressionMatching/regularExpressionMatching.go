package regularExpressionMatching

func isMatch(s string, p string) bool {
	const zeroOrMore = '*' // 42
	const oneChar = '.'    // 46

	dp := make([][]bool, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := i < len(s) && (p[j] == s[i] || p[j] == oneChar)

			if j+1 < len(p) && p[j+1] == zeroOrMore {
				dp[i][j] = dp[i][j+2] || firstMatch && dp[i+1][j]
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}

func isMatchShittyImplementation(s string, p string) bool {
	const zeroOrMore = '*' // 42
	const oneChar = '.'    // 46

	revS := reverse(s)
	revP := reverse(p)

	pi := 0
	si := 0
	var patternRune uint8
	var strRune uint8
	var runeToMatch uint8
	var firstTimeZeroOrMore bool
	var isDirty bool

	for {
		if pi <= len(revP)-1 {
			patternRune = revP[pi]
		}

		runeToMatch = patternRune

		if patternRune == zeroOrMore {
			if pi+1 >= len(revP) {
				return false
			}

			runeToMatch = revP[pi+1]

			if firstTimeZeroOrMore {
				if si < len(revS)-1 {
					si++
					isDirty = false
				} else {
					return true
				}
			}

			firstTimeZeroOrMore = true
		}

		if si <= len(revS)-1 {
			strRune = revS[si]
		}

		// =============================================================================================================
		// Comparing string with placeholders

		if strRune == runeToMatch || patternRune == oneChar || runeToMatch == oneChar {
			if patternRune != zeroOrMore && pi != len(revP) {
				pi++
				isDirty = true
			}

			if si != len(revS) {
				si++
				isDirty = false
			}
		} else if patternRune == zeroOrMore {
			pi = pi + 2
			isDirty = true
			firstTimeZeroOrMore = false
			continue
		} else {
			return false
		}

		if pi == len(revP) && si == len(revS) && !isDirty {
			return true
		}

		if pi == len(revP) {
			return false
		}

		// End of comparing
		// =============================================================================================================
	}
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
