package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"company/vpngo/server/common"
)

func main() {
	appId := "123"
	appSercet := "8elSYpCKwN"
	t := time.Now().Unix()
	str := fmt.Sprintf("%s%s%d", appSercet, appId, t)
	sig := common.SHA1Sign(str)
	getUrl := fmt.Sprintf("http://localhost:8867/api/v1/app/config/dns_info?app_id=%s&timestamp=%d&sign=%s", appId, t, sig)
	resp, err := http.Get(getUrl)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))
}
