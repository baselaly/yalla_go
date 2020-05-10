package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims We add as an embedded type in jwt token
type Claims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

// define jwt key
var jwtKey = []byte("my_secret_key")

// CreateToken create jwt token for user by id in claims
func CreateToken(ID uint) (string, error) {
	// expiration time
	expirationTime := time.Now().Add(100 * 100 * time.Minute)

	claims := Claims{
		ID: ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
