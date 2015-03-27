package tests

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"luck-web/utils"
	"strings"
	"testing"
)

func TestScrypt(t *testing.T) {
	password := "im lucky"
	salt := strings.Replace(uuid.NewUUID().String(), "-", "", -1)
	newkey := utils.GetKey(password, salt)
	fmt.Printf("pro :password %s  salt %s  = %s", password, salt, newkey)
}
