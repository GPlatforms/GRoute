package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
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

func (a *AesEncrypt) Encrypt(str []byte, secret string) (string, error) {
	key := a.validSecret(secret)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	origData := a.pKCS5Padding(str, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func (a *AesEncrypt) Decrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = a.pKCS5UnPadding(origData)

	return origData, nil
}

func (a *AesEncrypt) pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (a *AesEncrypt) pKCS5UnPadding(origData []byte) []byte {
	lenght := len(origData)
	unpadding := int(origData[lenght-1])
	return origData[:(lenght - unpadding)]
}

func (a *AesEncrypt) zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func (a *AesEncrypt) zeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
