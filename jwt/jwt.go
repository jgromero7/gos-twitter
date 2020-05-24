package jwt

import (
	"errors"
	"strings"

	"github.com/jgromero7/gos-twitter/app/models"
	"github.com/jgromero7/gos-twitter/app/structs"

	jwt "github.com/dgrijalva/jwt-go"
)

// UserID id current user Auth
var UserID string

// UserEmail email current user Auth
var UserEmail string

// GenerateJWT method generate one token for user
func GenerateJWT(payload jwt.MapClaims) (string, error) {

	privateKey := []byte("gos-twitter-app-private-key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}

// DecodeJWT docode token, get info
func DecodeJWT(token string) (*structs.Claim, bool, string, error) {
	privateKey := []byte("gos-twitter-app-private-key")

	claim := &structs.Claim{}

	auxToken := strings.Split(token, "Bearer")
	if len(auxToken) != 2 {
		return claim, false, "", errors.New("Format Token Ivalid")
	}

	token = strings.TrimSpace(auxToken[1])

	currentToken, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	if err == nil {
		_, exists, _ := models.ExistsUser(claim.Email)
		if exists == true {
			UserEmail = claim.Email
			UserID = claim.ID.Hex()
		}

		return claim, exists, UserID, nil
	}

	if !currentToken.Valid {
		return claim, false, "", errors.New("Format Token Ivalid")
	}

	return claim, false, "", err
}
