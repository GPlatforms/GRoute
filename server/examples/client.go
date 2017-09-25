package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"company/vpngo/server/common"
)

func main() {
	appId := "11235"
	appSercet := "8elSYpCKwN"
	secert := []byte("w2FG8DjogqMaM5Do")
	t := time.Now().Unix()
	str := fmt.Sprintf("%s%s%d", appSercet, appId, t)
	sig := common.SHA1Sign(str)
	getUrl := fmt.Sprintf("http://121.40.210.113/groute/v1/config?app_id=%s&timestamp=%d&sign=%s", appId, t, sig)
	resp, err := http.Get(getUrl)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))
	aesData, _ := base64.StdEncoding.DecodeString(string(b))

	aesEnc := common.AesEncrypt{}
	aes, err := aesEnc.Decrypt(aesData, secert)

	fmt.Println(string(aes), err)
}
