package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type SmartLabJwtClaims struct {
	echo.Context
	Token           string `json:"token"`
	ParentCompanyID int    `json:"parent_id"`
	UserID          int64  `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken create JWT token with encrypted claims
func GenerateToken(jsession string, parentCompanyID int, userID int64, validDuration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, SmartLabJwtClaims{
		Token:           jsession,
		UserID:          userID,
		ParentCompanyID: parentCompanyID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(validDuration).Unix(),
		},
	})

	hmacSampleSecret := []byte(viper.GetString("AES_PKEY_32BIT"))
	return token.SignedString(hmacSampleSecret)
}

func ParseToken(tokenString string) (string, int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if token.Method.Alg() != "HS256" {
			return nil, errors.New("UNEXPECTED SIGNING METHOD")
		}
		return []byte(viper.GetString("AES_PKEY_32BIT")), nil
	})
	if err != nil {
		return "", -1, err
	}
	if !token.Valid {
		return "", -1, errors.New("INVALiD TOKEN CHECK")
	}
	if token.Claims.Valid() != nil {
		return "", -1, errors.New("INVALiD TOKEN CLAIMS CHECK")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		strTok, okAssert := claims["token"].(string)
		if !okAssert {
			return "", -1, errors.New("INVALID TOKEN TYPE")
		}
		intPCompID, okokAssert := claims["parent_id"].(int)
		if !okokAssert {
			return "", -1, errors.New("INVALID TOKEN TWO TYPE")
		}
		return strTok, intPCompID, nil
	}
	return "", -1, errors.New("INVALID TOKEN")

}
