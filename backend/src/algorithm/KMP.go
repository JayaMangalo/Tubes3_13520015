package algorithm

func KMP(text, pattern string) int {
	var n int = len(text)
	var m int = len(pattern)

	var fail = computeFail(pattern)

	var i int = 0
	var j int = 0
	for i < n {
		if pattern[j] == text[i] {
			if j == m-1 {
				return i - m + 1
			}
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			i++
		}
	}
	return -1
}

func computeFail(pattern string) []int {
	var m int = len(pattern)

	var fail = make([]int, m)
	fail[0] = 0

	var j int = 0
	var i int = 1
	for i < m {

		if pattern[j] == pattern[i] {
			fail[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			fail[i] = 0
			i++
		}
	}
	return fail
}
