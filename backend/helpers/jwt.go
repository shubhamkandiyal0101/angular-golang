package jwttokenhelper

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte("very_secret_jwt_secret_key")

// Tutorial 01: https://www.codershood.info/2020/03/14/jwt-authentication-in-golang-tutorial-with-example-api/
// Tutorial 02: https://dev.to/joojodontoh/build-user-authentication-in-golang-with-jwt-and-mongodb-2igd

// CreateJWT func will used to create the JWT while signing in and signing out
func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(jwtSecretKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

// VerifyToken func will used to Verify the JWT Token while using APIS
func VerifyToken(tokenString string) (email string, err error) {
	tokenData, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return jwtSecretKey, nil
	})

	if tokenData.Valid {
		if tokenClaims, ok := tokenData.Claims.(jwt.MapClaims); ok && tokenData.Valid {
			return fmt.Sprint(tokenClaims["email"]), nil
		} else {
			var customErr error
			customErr = errors.New("Invalid Token")
			return "", customErr
		}
	} else {
		var customErr error
		customErr = errors.New("Invalid Token")
		return "", customErr
	}

}
