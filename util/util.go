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

//
func ReadCookie(cookie_key string, req *http.Request) {

// readCookies parses all "Cookie" values from resquest
	res.Cookies()
	
	lines, ok := h["Cookie"]
	if !ok {
		return []*Cookie{}
	}

	cookies := []*Cookie{}
	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), ";")
		if len(parts) == 1 && parts[0] == "" {
			continue
		}
		// Per-line attributes
		for i := 0; i < len(parts); i++ {
			parts[i] = strings.TrimSpace(parts[i])
			if len(parts[i]) == 0 {
				continue
			}
			name, val := parts[i], ""
			if j := strings.Index(name, "="); j >= 0 {
				name, val = name[:j], name[j+1:]
			}
			if !isCookieNameValid(name) {
				continue
			}
			if filter != "" && filter != name {
				continue
			}
			val, ok := parseCookieValue(val, true)
			if !ok {
				continue
			}
			cookies = append(cookies, &Cookie{Name: name, Value: val})
		}
	}
	return cookies
}
///
	c, err := req.Cookie(cookie_key)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c.Value)
}
