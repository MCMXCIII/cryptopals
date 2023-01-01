package main

import (
	"fmt"
)

func decrypt(ciphertext string) {
	for key :=0; key < 256; key++ {
		plaintext := ""
		//for i, c := range ciphertext {// I inst used.
		for _, c := range ciphertext {
			plaintext += string(key ^ int(c))
		}
		fmt.Printf("Key  %d Plaintext: %s\n", key, plaintext)
	}
}

func main() {
	ciphertext := 
	"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decrypt(ciphertext)
}

