package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"company/vpngo/server/common"
)

func main() {
	appId := "11235"
	appSercet := "8elSYpCKwN"
	t := time.Now().Unix()
	str := fmt.Sprintf("%s%s%d", appSercet, appId, t)
	sig := common.SHA1Sign(str)
	getUrl := fmt.Sprintf("http://localhost:8867/groute/v1/config?app_id=%s&timestamp=%d&sign=%s", appId, t, sig)
	resp, err := http.Get(getUrl)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))
}
