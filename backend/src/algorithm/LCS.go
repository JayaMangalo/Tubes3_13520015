package algorithm

func LCSPercentage(text, pattern string) int {
	same := LCS(text, pattern,len(text),len(pattern),0)
	percentage := same * 100 / len(pattern)
	return percentage
}

// func LCS(text, pattern string) int {
// 	m := len(text)
// 	n := len(pattern)

// 	prev := make([]int, n+1)
// 	cur := make([]int, n+1)
// 	temp := make([]int, n+1)

// 	for i := 1; i < m+1; i++ {
// 		for j := 1; j < n+1; j++ {
// 			if text[i-1] == pattern[j-1] {
// 				cur[j] = 1 + prev[j-1]
// 			} else {
// 				cur[j] = max(cur[j-1], prev[j])
// 			}
// 		}
// 		copy(temp, cur)
// 		copy(cur, prev)
// 		copy(prev, temp)
// 	}
// 	return prev[n]
// }

// func main(){
// 	s1 := "GXTXAYB"
// 	s2 := "AGGTAB"
// 	print("Length of LCS is ", LCS(s1, s2))
// 	print("Similarity is ", LCSPercentage(s1, s2))
// }
func LCS(X, Y string, i,j,count int) int{
	if (i == 0 || j == 0){
		return count;
	}

	if (X[i - 1] == Y[j - 1]){
		count = LCS(X,Y,i - 1, j - 1, count + 1);
	}
	count = max(count,max(LCS(X,Y, i, j - 1, 0),LCS(X,Y,i - 1, j, 0)));
	return count;
    }