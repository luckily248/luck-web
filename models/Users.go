package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"luck-web/utils"
)

type User struct {
	BaseDBmodel
	Email    string
	Password string
	Salt     string
}

func (this *User) Tablename() string {
	return "users"
}
func (this *User) init() {
	this.BaseDBmodel.init()
	this.c = this.db.C(this.Tablename())
}

func AddUser(user User) error {
	user.init()
	defer user.session.Close()
	return user.c.Insert(user)
}
func CheckUser(email string, password string) (bool, string) {
	user := User{}
	user.init()
	defer user.session.Close()
	fmt.Printf("start find %s\n", email)
	err := user.c.Find(bson.M{"email": email}).One(&user)
	fmt.Printf("email,%s,salt,%s,passowrd,%s\n", user.Email, user.Salt, user.Password)
	if err != nil {
		fmt.Println("find err")
		return false, ""
	}
	npassword := utils.GetKey(password, user.Salt)
	if user.Password != npassword {
		fmt.Println("not match")
		return false, ""
	}

	return true, npassword
}
