package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	plain := "Hello World!"
	fmt.Println("[ Encoding  ]")
	str := base64.StdEncoding.EncodeToString([]byte(plain))
	fmt.Println("Plain: ", plain)
	fmt.Println("Base64: ", str)
	fmt.Println("[ Decoding  ]")
	fmt.Println("Base64: ", str)
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)

	}
	fmt.Printf("Plain:\t%s\n", data)

}
