package until

import (
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// 加密密码
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		beego.Error("")
	}
	return string(hash)
}

// 验证密码
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
