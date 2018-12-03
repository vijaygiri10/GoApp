package util

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtSigningScrect string = "maropost_jwt_screct"

//GenrateJWT check jwt valid or not
func GenrateJWT() (string, error) {

	//Create a new token object, specifying signing method
	token := jwt.New(jwt.SigningMethodHS256)

	//Sign and get the complete encoded token as a string using the secret
	return token.SignedString(JwtSigningScrect)
}

//ValidateJWT check jwt valid or not
func ValidateJWT(res *http.Request) (bool, error) {

	//Grab the token from the header
	tokenHeader := res.Header.Get("Authorization")

	//Check Token is missing or not
	if tokenHeader == "" {
		return false, errors.New("Missing auth token")
	}

	//The token normally comes in format `Bearer {token-body}`,
	//we check if the retrieved token matched this requirement
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		return false, errors.New("Invalid/Malformed auth token")
	}

	//Grab the token part, what we are truly interested in
	jwtPart := splitted[1]

	token, err := jwt.Parse(jwtPart, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// your secret maropost_jwt_screct
		return []byte(JwtSigningScrect), nil
	})

	//Malformed token, returns with http code 403 as usual
	if err != nil {
		return false, errors.New("Malformed authentication token")
	}

	//Token is invalid, maybe not signed on this server
	if !token.Valid {
		return false, errors.New("Token is not valid.")
	}
	return true, nil
}

//GetAllCookies from http request
func GetAllCookies(req *http.Request) (cookieMap map[string]string) {

	cookieMap = map[string]string{}

	for _, cookie := range req.Cookies() {
		cookieMap[cookie.Name] = cookie.Value
	}

	return cookieMap
}

//GetCookieByName from http request
func GetCookieByName(By_Name string, req *http.Request) (string, error) {

	cookieObject, err := req.Cookie(By_Name)
	if err != nil {
		return "", err
	}

	return cookieObject.Value, nil
}
