package regularExpressionMatching

func isMatch(s string, p string) bool {
	const oneChar = '.'

	ji := 0
	for i := 0; i < len(p); i++ {
		patternRune := p[i]

		if ji == len(s) {
			return false
		}

	strLoop:
		for j := ji; j < len(s); j++ {
			strRune := s[j]

			switch true {
			case patternRune == oneChar:
				ji = j + 1
				break strLoop
			case patternRune == strRune:
				ji = j + 1
				break strLoop
			default:
				return false
			}
		}
	}

	return true
}
