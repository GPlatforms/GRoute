package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/hex"
)

func SHA1Sign(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

type AesEncrypt struct{}

func (a *AesEncrypt) validSecret(secret string) []byte {
	lenght := len(secret)
	if lenght < 16 {
		panic("res key 长度不能小于16")
	}
	return []byte(secret)[:16]
}

func (a *AesEncrypt) Encrypt(str []byte, secret string) ([]byte, error) {
	key := a.validSecret(secret)
	iv := key[:aes.BlockSize]

	encrypted := make([]byte, len(str))

	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, str)
	return encrypted, nil
}
