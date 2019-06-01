package service

import (
	"Grocery-Shopping-Token-Management-Module/app/handler"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)
// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Credentials struct {
	username string `json:"username"`
	password string `json:"password"`
}

func GenerateFunction(w http.ResponseWriter,r * http.Request) {
	var creds  Credentials
	expirationTime := time.Now().Add(5 * time.Minute)
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		handler.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	claims := Claims{
		Username: creds.username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusBadRequest)
		handler.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}
	handler.RespondJSON(w,http.StatusOK,tokenString)
}