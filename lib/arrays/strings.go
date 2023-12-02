package arrays

// Splits string into a flat string array with any number of split strings.
func SplitWithAll(str string, match []string) []string {
	// order `match` to be from longest to shortest
	if len(match) == 0 {
		return []string{str}
	}
	var res []string
	var lastI int

	for i := 0; i < len(str); i++ {
		for _, pattern := range match {
			if i+len(pattern) > len(str) {
				continue
			}

			if str[i:i+len(pattern)] != pattern {
				continue
			}

			if len(str[lastI:i]) != 0 {
				res = append(res, str[lastI:i])
			}

			if len(pattern) != 0 {
				res = append(res, pattern)
			}
			lastI = i + len(pattern)
			i += len(pattern) - 1
			break
		}
	}

	if len(str[lastI:]) != 0 {
		res = append(res, str[lastI:])
	}

	return res
}
