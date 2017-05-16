package utils

import (
	"fmt"
	"github.com/elithrar/simple-scrypt"
	"log"
)

// Hash create and return hash from given password
func Hash(password []byte) string {
	hash, err := scrypt.GenerateFromPassword(password, scrypt.DefaultParams)

	fmt.Println(hash)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
