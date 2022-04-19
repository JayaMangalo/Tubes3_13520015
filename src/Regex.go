package main

import (
    "regexp"
)
  
func isSanitizedRegex(text string) bool{
	match1, err := regexp.MatchString("^[ATCG]+$", text)
	if err == nil{
		return match1
	}
	return false
}

// func main() {
// 	if isSanitizedRegex("ATTCGTAACTAGTAAGTTA"){
// 		print("yes")
// 	}else{
// 		print("no")
// 	}
// }