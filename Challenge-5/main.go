package main 

import (
	"encoding/hex"
	"fmt"
)

func main() {
	//the plaintext to use
	plaintext := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	//the Key to use for Repeating XOR
	key := "ICE"
	
	//convert the plaintext message to a slice of bytes
	plaintextBytes := []byte(plaintext)

	//Initilize a slice to store the message
	ciphertext := make([]byte, len(plaintextBytes))

	//Perform repeating-key XOR
	for i, b := range plaintextBytes {
		ciphertext[i] = b ^ key [i % len(key)]
	}

	//Print the ciphertext as a hex string
	fmt.Println(hex.EncodeToString(ciphertext))
}
