package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	rand2 "math/rand"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/juliotorresmoreno/freelive/configs"
	"github.com/juliotorresmoreno/freelive/model"
	"github.com/juliotorresmoreno/freelive/services"
	"github.com/labstack/echo/v4"
)

func GetAesKey(secret string) []byte {
	conf := configs.GetConfig()
	s := []byte(conf.Secret)[:16]
	h := md5.New()
	_, err := io.WriteString(h, secret)

	if err != nil {
		log.Fatal(err)
	}

	seed := binary.BigEndian.Uint64(h.Sum(nil))
	r := rand2.New(rand2.NewSource(int64(seed)))
	t := fmt.Sprintf("%v%v", r.Int31(), r.Int31())[:16]

	return append(s, t...)
}

func Encrypt(key []byte, message string) (string, error) {
	plainText := []byte(message)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.URLEncoding.EncodeToString(cipherText), nil
}

func Decrypt(key []byte, securemess string) (string, error) {
	cipherText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("Ciphertext block size is too short")
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func MakeSession(c echo.Context, u *model.User) error {
	token := bson.NewObjectId().Hex()
	redisCli := services.NewRedis()
	redisCli.Set(token, u.ID, 24*time.Hour)
	go redisCli.Close()

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   60 * 60 * 10,
		HttpOnly: true,
	})

	return c.JSON(200, u)
}
