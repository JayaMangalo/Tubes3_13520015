package main

import (
	"main/src/handler"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

// import "fmt"

var jwtMiddleware *jwtmiddleware.JWTMiddleware

func main() {

	/*
		text, err := ReadFile("disease/Covid.txt")
		text2, err2 := ReadFile("people/Gerald.txt")

		if err == false || err2 == false {
			return
		}
		if isSanitizedRegex(text) == false || isSanitizedRegex(text2) == false {
			return
		}

		var index int = LCSPercentage(text2, text)
		println(index)
	*/

	// TES NGAMBIL DATA DARI DATABASE
	//connector.GetDataOrang()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowedMethods:  []string{"GET", "POST"},
	}))
	router.GET("/diagnosis/:query", handler.DiagnosisHandler)
	router.POST("/inputPenyakit/:nama_penyakit/:dna_sequence", handler.PostPenyakitHandler)
	router.POST("/tesDNA/:nama_pengguna/:nama_penyakit/:dna_sequence", handler.TesDNA)

	// Start and run the server
	router.Run("localhost:8080")

}
