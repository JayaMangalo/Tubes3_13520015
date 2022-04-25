package algorithm

func BoyerMoore(text, pattern string) int {
	var last [128]int = buildLast(pattern)
	var n int = len(text)
	var m int = len(pattern)
	var i = m - 1

	if i > n-1 {
		return -1
	}

	var j = m - 1

	for {
		if pattern[j] == text[i] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			var lo int = last[text[i]]
			i = i + m - min(j, i+lo)
			j = m - 1
		}
		if i > n-1 {
			break
		}
	}
	return -1
}

func buildLast(pattern string) [128]int {
	var last [128]int
	for i := 0; i < 128; i++ {
		last[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		last[pattern[i]] = i
	}
	return last
}
