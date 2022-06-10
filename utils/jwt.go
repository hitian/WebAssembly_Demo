package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func JwtGenerate(secret, content string) (string, error) {
	//convert content to jwt.MapClaims ()
	value := make(map[string]interface{})
	if err := json.Unmarshal([]byte(content), &value); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(value))
	signedString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("err: %s, top level key must string", err)
	}

	return signedString, nil
}

func JwtVerify(secret, content string) (bool, string, error) {
	token, _ := jwt.Parse(content, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token check failed")
		}
		return []byte(secret), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, "", errors.New("content decode failed. not jwt token?")
	}

	value, err := json.Marshal(claims)
	if err != nil {
		return token.Valid, "", err
	}

	return token.Valid, string(value), nil
}
