package util

import (
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtSigningKey string = "maropost_jwt_screct"

//GenrateJWT check jwt valid or not
func GenrateJWT() {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

}

//ValidateJWT check jwt valid or not
func ValidateJWT(res *http.Request) (bool, error) {

	//Grab the token from the header
	tokenHeader := res.Header.Get("Authorization")

	//Token is missing, returns with error code 403 Unauthorized
	if tokenHeader == "" {
		return false, errors.New("missing auth token")
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
		return []byte("maropost_jwt_screct"), nil
	})

	//Malformed token, returns with http code 403 as usual
	if err != nil {
		return false, errors.New("Malformed authentication token")
	}

	//Token is invalid, maybe not signed on this server
	if !token.Valid {
		return false, errors.New("Token is not valid.")
	}
}

//GetAllCookies
func GetAllCookies(req *http.Request) (cookieMap map[string]string) {

	cookieMap = map[string]string{}

	for _, cookie := range req.Cookies() {
		cookieMap[cookie.Name] = cookie.Value
	}

	return cookieMap
}

//GetCookieByName
func GetCookieByName(By_Name string, req *http.Request) string {

	cookieObject, err := req.Cookie(By_Name)
	if err != nil {
		return ""
	}

	return cookieObject.Value
}
