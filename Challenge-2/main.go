package main


import "fmt"
import "encoding/hex"

func xorBuffers(buf1, buf2 []byte) []byte{
	// check that buffers are the same length
	if len (buf1) != len(buf2) {
		panic("Buffers must be the same length!")
	}

	//Init the result buffer
	result := make([]byte , len(buf1))

	//Xor the buffers together
	for i, b1 := range buf1 {
		result[i] = b1 ^ buf2[i]
	}

	return result
}

// A function used the test Buffers
//func main() {
	//parse the input of the string
//	buf1 := []byte{0x01,0x02,0x03,0x04}
//	buf2 := []byte{0x01,0x02,0x03,0x04}
//	result := xorBuffers(buf1, buf2)
//	fmt.Println(result) // should print [4 4 4 4]
//}

func main(){
// Parse the input buffers from the hex strings
	buf1, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		panic(err)
	}
// Use the XORing string to compare "686974207468652062756c6c277320657965"
	buf2, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		panic(err)
	}

	//XOR the buffers and print the result
	result := xorBuffers(buf1, buf2)
	fmt.Printf("%x\n", result) // prints the 746865206b696420646f6e277420706c6179 string expected 
//	fmt.Printf("%x\n", result) // Prints "The kid don't play"
}
