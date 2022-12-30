package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	//Convert the String to a byte slice
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hexBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Encode the slice to a base 64 string
	base64Str := base64.StdEncoding.EncodeToString(hexBytes)

	//Print the string now
	fmt.Println(base64Str)
}
