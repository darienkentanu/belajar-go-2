package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "john wick"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded:", decodedString)

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString2 = string(encoded)
	fmt.Println(encodedString2)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString2 = string(decoded)
	fmt.Println(decodedString2)

	data = "https://kalipare.com/"
	encodedString3 := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString3)

	decodedByte2, _ := base64.URLEncoding.DecodeString(encodedString3)
	decodedString = string(decodedByte2)
	fmt.Println(decodedString)
}
