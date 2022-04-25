package main

import (
	"main/src/handler"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/gin-gonic/contrib/static"
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

	router.Use(static.Serve("/", static.LocalFile("./front-end/public", true)))
	router.GET("/diagnosis", handler.DiagnosisHandler)
	router.POST("/inputPenyakit/:nama_penyakit/:dna_sequence", handler.PostPenyakitHandler)
	router.GET("/tesDNA/:nama_pengguna/:nama_penyakit/:dna_sequence", handler.TesDNA)

	// Start and run the server
	router.Run("localhost:3000")

}

// YANG DIBAWAH KOMEN INI BUAT PROTEKSI API
// 1. gatau perlu ato ga
// 2. belum ngerti cara kerjannya
/*
func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(os.Getenv("AUTH0_DOMAIN") + ".well-known/jwks.json")
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	x5c := jwks.Keys[0].X5c
	for k, v := range x5c {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + v + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return cert, errors.New("unable to find appropriate key")
	}

	return cert, nil
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client secret key
		err := jwtMiddleware.CheckJWT(c.Handler)
		if err != nil {
			// Token not found
			fmt.Println(err)
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Unauthorized"))
			return
		}
	}
}*/

/*
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			aud := os.Getenv("AUTH0_API_AUDIENCE")
			checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAudience {
				return token, errors.New("Invalid audience.")
			}
			// verify iss claim
			iss := os.Getenv("AUTH0_DOMAIN")
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			cert, err := getPemCert(token)
			if err != nil {
				log.Fatalf("could not get cert: %+v", err)
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	// register our actual jwtMiddleware
	jwtMiddleWare := jwtMiddleware
*/

// ISI FILE .env:
/*export AUTH0_API_CLIENT_SECRET="Gi-MHv1BQDQNnN0ZLcb-NTkWVqCKWa7AXMdQIXmQg70Nb0uTUjIaaNg-4EZYEAiT"
export AUTH0_CLIENT_ID="9AK5eY2EYVC5U4kIIVbv9QF1LnmAvkGv"
export AUTH0_DOMAIN="dev-0ki6e9pj.us.auth0.com"
export AUTH0_API_AUDIENCE=""*/
