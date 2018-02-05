package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)


	sDec := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(sDec)

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec,_ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}