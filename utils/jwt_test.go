package utils

import (
	"fmt"
	"testing"
)

func TestJWTGenerate(t *testing.T) {
	secret := "Yt1Kq7n5kf24eSGg9"
	content := `{"user": 123}`

	signedString, err := JwtGenerate(secret, content)
	if err != nil {
		fmt.Println("err:", err)
		t.FailNow()
	}

	fmt.Println(signedString)
}

func TestJWTVerify(t *testing.T) {
	secret := "Yt1Kq7n5kf24eSGg9"
	content := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoxMjN9.evcrRAFM5ZBP62KA9oE8gT9vFbg_5dQgp7_J8Q9WUhw`

	ok, json, err := JwtVerify(secret, content)
	if err != nil && ok {
		fmt.Println("err:", err)
		t.FailNow()
	}
	fmt.Println(json)
}
