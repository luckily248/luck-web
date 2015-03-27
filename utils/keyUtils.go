package utils

import (
	"fmt"
	"github.com/dchest/scrypt"
)

//scrypt加密
func GetKey(password string, salt string) string {
	newkey, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	return fmt.Sprintf("%x", newkey)
}
