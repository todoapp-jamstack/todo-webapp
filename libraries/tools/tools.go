package tools

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt is used to hash and salt a string
// it has to be already a byte tho
func HashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// CreateFolder is used to create folders
func CreateFolder(path string) bool {

	// create db folder
	err := os.Mkdir(path, 0755)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}
