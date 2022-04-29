package algorithm

//fungsi bantuan yang digunakan untuk mencari kemiripan pattern 
//dengan text dalam bentuk integer persentase yang direturn
func LCSPercentage(text, pattern string) int {
	same := LCS(text, pattern)
	percentage := same * 100 / len(pattern)
	return percentage
}
	
//fungsi LCS untuk mencari least common substring
func LCS(text,pattern string) int {
	m := len(text)
	n := len(pattern)
	L := make([][]int, m+1)   
	
	var length int 
	length = 0
	for i := 0; i < m+1; i ++ {
		L[i] = make([]int, n+1)  
			for j := 0; j < n+1; j++ {
				if i == 0 || j == 0 {
					L[i][j] = 0;
					} else if text[i-1] == pattern[j-1] {
						L[i][j] = L[i-1][j-1] + 1
						length = max(L[i][j],length);
					} else {
				L[i][j] = 0;
			}
		}
	}
	
	return length
	}