package main

import (
	"fmt"
	"testing"

	jsonhelper "example.com/webassembly/cmd/wasm/json"
	"example.com/webassembly/cmd/wasm/qrcode"
)

func TestPreetyJson(t *testing.T) {
	const rawJson = `{"website":"golangbot.com", "tutorials": {"string":"https://golangbot.com/strings/"}}`
	pretty, err := jsonhelper.PrettyJson(rawJson)
	if err != nil {
		fmt.Println("error occur!", err)
	}
	fmt.Println(pretty)
}
func TestGenerateQr(t *testing.T) {
	//firstName := "lan"
	//lastName := "tian"
	//mail := "lant@neusoft.com"
	//phone := "15668668395"
	//qrcode.GenerateQr(firstName, lastName, mail, phone)
	if base64Str, err := qrcode.GenerateQr("http://www.baidu.com"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(base64Str)
	}

}
