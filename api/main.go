package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

var SigningKey = []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Super Secret Information")
}

// Middleware
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf(("Invalid signing method"))
				}

				aud := "billing.jwtgo.io"
				checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)

				if !checkAudience {
					return nil, fmt.Errorf(("Invalid aud"))
				}

				iss := "jwtgo.io"
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)

				if !checkIss {
					return nil, fmt.Errorf(("Invalid iss"))
				}

				return SigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}	
		} else {
			fmt.Fprintf(w, "No authorization token provided")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Printf("Server is ready to check autharization. Listening on port: 9001\nExample: curl http://localhost:9001 --header 'Token: ***'")
	handleRequests()
}