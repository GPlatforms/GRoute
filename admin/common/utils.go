package common

import (
	cryrand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	// "math"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JWT(priKey []byte, username string, tokenExpire int64) string {
	claims := MyCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(tokenExpire)).Unix(),
			Issuer:    "vpn",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(priKey)
	return tokenString
}

func ParseJWT(priKey []byte, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return priKey, nil
	})
}

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func GenRsaKey(bits int, priFileName string) error {
	//如果私钥不存在则生成密钥，否则不处理
	if !FileExist(priFileName) {
		// 生成私钥文件
		privateKey, err := rsa.GenerateKey(cryrand.Reader, bits)
		if err != nil {
			return err
		}
		derStream := x509.MarshalPKCS1PrivateKey(privateKey)
		block := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: derStream,
		}
		file, err := os.Create(priFileName)
		if err != nil {
			return err
		}
		err = pem.Encode(file, block)
		if err != nil {
			return err
		}
	}
	return nil
}

// func PageHandle(page, count int) []int {
// 	x := int(math.Ceil(float64(count) / 20))

// 	if page < 5 {

// 	} else {
// 		if page <= x {

// 		}
// 	}
// }
