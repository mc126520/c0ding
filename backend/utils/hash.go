package utils

import "golang.org/x/crypto/bcrypt"

//对密码进行加密
func HashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashpassword), err
}

//检查密码
func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
