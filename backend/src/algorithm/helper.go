package algorithm

import(
	"strings"
)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkifDate(text string) bool{
	arr := [12]string{"JANUARY","FEBRUARY","MARCH","APRIL","MAY","JUNE","JULY","AUGUST","SEPTEMBER","OCTOBER","NOVEMBER","DECEMBER"}
	var text2 = strings.ToUpper(text);

	for _, x := range arr {
        if x == text2 {
            return true
        }
    }
	return false
}
