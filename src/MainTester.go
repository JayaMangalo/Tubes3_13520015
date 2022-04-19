package main
// import "fmt"

func sanitize(){}


func main() {
	text,err := ReadFile("disease/Covid.txt")
	text2,err2 := ReadFile("people/Gerald.txt")
	
	if(err == false || err2 == false){
		return
	}
	if(isSanitizedRegex(text) == false || isSanitizedRegex(text2) == false ){
		return
	}

	var index int = LCSPercentage(text2,text)
	println(index)
}