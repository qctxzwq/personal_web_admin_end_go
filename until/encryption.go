package until

import (
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		beego.Error("password hash error:", err)
	}
	return string(hash)
}

// 验证密码
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		beego.Error("password compare error:", err)
		return false
	}
	return true
}
