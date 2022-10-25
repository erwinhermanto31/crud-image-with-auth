package utils

import (
	"encoding/base64"
	"log"
	"strconv"
	"time"

	"github.com/erwinhermanto31/crud-image-with-auth/utils/errors"
	"github.com/golang-jwt/jwt"
)

const (
	DateFormatRFC3339 = time.RFC3339
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	ID       int32  `json:"id"`
	jwt.StandardClaims
}

type ClaimsRes struct {
	Username string `json:"username"`
	ID       int32  `json:"id"`
}

func FormatDateToRFC3339(t time.Time) string {
	return t.Format(DateFormatRFC3339)
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func GenerateToken(username string, id int32, expiresAt int64, subject string) (tokenString string) {
	// Create the JWT claims, which includes the username and expiry time
	log.Printf("[GenerateToken] request : %v", username)
	log.Printf("[GenerateToken] request : %v", id)
	claims := &Claims{
		Username: username,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expiresAt,
			Subject:   subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ = token.SignedString(jwtKey)

	return tokenString
}

func ParsingToken(tokenString string) (claimsRes ClaimsRes, err error) {
	// mapString := make(map[string]string)

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		log.Printf("[ParsingToken] error get : %v", err)
		return claimsRes, err
	}

	if !tkn.Valid {
		log.Printf("[ParsingToken] token not valid : %v", err)
		return claimsRes, errors.ErrNotFound("token not valid")
	}

	claimsRes.Username = claims.Username
	claimsRes.ID = claims.ID

	log.Printf("[ParsingToken] response : %v", claims)
	log.Printf("[ParsingToken] response : %v", claimsRes)
	return claimsRes, nil
}
