package regularExpressionMatching

//"mississippi", // 3, second 46
//"mis*is*Ip*.", // 3-4

func isMatch(s string, p string) bool {
	const zeroOrMore = '*' // 42
	const oneChar = '.'    // 46

	revS := reverse(s)
	revP := reverse(p)

	pi := 0
	si := 0
	var patternRune uint8
	var strRune uint8
	var runeToMatch uint8

	for {
		if pi < len(revP)-1 {
			patternRune = revP[pi]
		}

		if si < len(revS)-1 {
			strRune = revS[si]
		}

		runeToMatch = patternRune

		if patternRune == zeroOrMore {
			runeToMatch = p[pi+1]

			if si < len(revS)-1 {
				// TODO this skips first char
				si++
			} else {
				return true
			}
		}

		switch {
		case strRune == runeToMatch || patternRune == oneChar || runeToMatch == oneChar:
			if patternRune != zeroOrMore {
				pi++
			}

			si++

			if pi == len(revP) && si == len(revS) {
				return true
			}
		default:
			if patternRune == zeroOrMore {
				pi++
				continue
			}

			return false
		}
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
