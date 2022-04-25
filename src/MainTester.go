package main

import (
	// "log"
	// "main/src/handler"
	// "net/http"
)

// import "fmt"

func sanitize() {}

// func main() {
// 	/*
// 		text, err := ReadFile("disease/Covid.txt")
// 		text2, err2 := ReadFile("people/Gerald.txt")

// 		if err == false || err2 == false {
// 			return
// 		}
// 		if isSanitizedRegex(text) == false || isSanitizedRegex(text2) == false {
// 			return
// 		}

// 		var index int = LCSPercentage(text2, text)
// 		println(index)
// 	*/

// 	// TES NGAMBIL DATA DARI DATABASE
// 	//connector.GetDataOrang()

// 	mux := http.NewServeMux()

// 	// handle route homepage(root)
// 	mux.HandleFunc("/", handler.HomeHandler)
// 	mux.HandleFunc("/inputPenyakit", handler.GetInputPenyakitHandler)
// 	mux.HandleFunc("/processInput", handler.PostPenyakitHandler)
// 	mux.HandleFunc("/lihatDiagnosis", handler.DiagnosisHandler)

// 	log.Println("Starting web on port 8080")

// 	err5 := http.ListenAndServe(":8080", mux)
// 	log.Fatal(err5)

// }
