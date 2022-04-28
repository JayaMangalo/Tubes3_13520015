package algorithm

//fungsi bantuan yang digunakan untuk mencari kemiripan pattern 
//dengan text dalam bentuk integer persentase yang direturn
func LCSPercentage(text, pattern string) int {
	same := LCS(text, pattern,len(text),len(pattern),0)
	percentage := same * 100 / len(pattern)
	return percentage
}

//fungsi rekursif Least Common Substring, me return panjang substring terbesar
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