package main

import (
	"fmt"
	"strings"
)

// Decrypt the ciphertext with the given key
func decrypt(ciphertext string, key int) string {
	plaintext := ""
	for _, c := range ciphertext {
		plaintext += string(key ^ int(c))
	}
	return plaintext
}

// Score the plaintext using character frequency analysis
func score(plaintext string) float64 {
	// Frequency of letters in English language
	frequencies := map[rune]float64{
		'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253,
		'e': 0.12702, 'f': 0.02228, 'g': 0.02015, 'h': 0.06094,
		'i': 0.06966, 'j': 0.00153, 'k': 0.00772, 'l': 0.04025,
		'm': 0.02406, 'n': 0.06749, 'o': 0.07507, 'p': 0.01929,
		'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056,
		'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150,
		'y': 0.01974, 'z': 0.00074,
	}

	// Convert plaintext to lowercase
	plaintext = strings.ToLower(plaintext)

	// Initialize score
	score := 0.0

	// Add the log of the frequency of each letter to the score
	for _, c := range plaintext {
		score += frequencies[c]
	}

	return score
}

func main() {
	ciphertext := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// Initialize best score and plaintext
	bestScore := 0.0
	bestPlaintext := ""
	bestKey := 0

	// Try decrypting the ciphertext with every possible key
	for key := 0; key < 256; key++ {
		plaintext := decrypt(ciphertext, key)
		score := score(plaintext)
		if score > bestScore {
			bestScore = score
			bestPlaintext = plaintext
			bestKey = key
		}
	}

	fmt.Printf("Key: %d  Plaintext: %s\n", bestKey, bestPlaintext)
}

