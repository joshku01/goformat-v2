package helper

import (
	"crypto/md5"
	"fmt"
	"goformat-v2/app/global/errorcode"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Md5Encryption md5加密
func Md5Encryption(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)

	return md5Str
}

// Md5EncryptionWithTime md5 加密（加上奈秒時間）
func Md5EncryptionWithTime(str string) string {
	naTime := time.Now().UnixNano()
	data := str + strconv.FormatInt(naTime, 10)
	key := []byte(data)

	token := md5.Sum(key)
	md5Str := fmt.Sprintf("%x", token)

	return md5Str
}

// HashPassword 密碼加密(註冊管理者使用)
func HashPassword(password string) (value string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		err = errorcode.ErrorHandler("ENCRYPT_FAILED", err)
		return string(bytes), err
	}

	return string(bytes), err
}

// CheckPasswordHash 檢查檢查(登入使用))
func CheckPasswordHash(password, dbPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(password))
	if err != nil {
		return false
	}
	return true
}

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset  中文編碼
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// StringRandom 字串亂數
func StringRandom(length int) string {
	return StringWithCharset(length, charset)
}
