package requests

import (
	"crypto/sha256"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Users struct {
	Login string `sql:"type varchar(50);not null;"`
	Pass  []byte `sql:"type varchar(50);not null;"`
	gorm.Model
}

type UserForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Login(login, password string, Db *gorm.DB) any {

	log.Printf("PARAMS LOGIN U:%v P:%v", login, password)

	if len(login) == 0 || len(password) == 0 {
		log.Printf("login empty params")
		return []byte("")
	}

	var user []Users
	res := Db.Find(&user, fmt.Sprintf("login = %v", login))
	if res.Error != nil {
		log.Printf("%v: login not found", res.Error)
		return []byte("")
	}

	hash := sha256.Sum256([]byte(password))

	log.Printf("\nres : %v\nhash: %v", user, hash)
	//if isEq(hash, )

	return nil
}

func Registration(login, password string, Db *gorm.DB) any {

	log.Printf("PARAMS REGISTRATION U:%v P:%v", login, password)

	if len(login) == 0 || len(password) == 0 {
		log.Printf("registration empty params")
		return []byte("")
	}

	res := Db.Find(&Users{}, login)
	if res.Error == nil {
		log.Printf("overwrite login")
		return []byte("")
	}
	hash := sha256.Sum256([]byte(password))
	regStr := Users{Login: login, Pass: hash[:]}

	res = Db.Create(&regStr)
	if res.Error != nil {
		log.Printf("%v: reg write db login", res.Error)
		return []byte("")
	}

	return fmt.Sprintf("login %v write", login)
}

func isEq(left, right []byte) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
