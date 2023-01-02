package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
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
		// Read the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Initialize best score and plaintext
	bestScore := 0.0
	bestPlaintext := ""
	bestKey := 0

	// Iterate over each line in the file
	for scanner.Scan() {
		// Convert the hex-encoded string to a byte slice
		input, err := hex.DecodeString(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		// Try decrypting the ciphertext with every possible key
		for key := 0; key < 256; key++ {
			plaintext := decrypt(string(input), key)
			score := score(plaintext)
			if score > bestScore {
				bestScore = score
				bestPlaintext = plaintext
				bestKey = key
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Key: %d  Plaintext: %s\n", bestKey, bestPlaintext)
}


