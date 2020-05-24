package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

// ExtractToken to extract token from header
func ExtractToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyToken to verify token
func VerifyToken(c *gin.Context) (uint, error) {
	// Get the JWT string from the header
	tokenString := ExtractToken(c)

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, err
		}
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("unauthorized")
	}

	return claims.ID, nil
}
