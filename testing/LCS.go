package main

func LCSPercentage(text, pattern string) int {
	same := LCS(text, pattern,len(text),len(pattern),0)
	percentage := same * 100 / len(pattern)
	return percentage
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
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

	func main(){
	s1 := "BXXXBAXXXCXXA"
	s2 := "BACA"
	print("Length of LCS is ", LCS(s1, s2,len(s1),len(s2),0))
	print("Similarity is ", LCSPercentage(s1, s2))
}