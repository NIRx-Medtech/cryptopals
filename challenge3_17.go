package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/robquant/cryptopals/pkg/tools"
)

var KEY, IV []byte

const BS = 16

var s string

var inputStrings = []string{
	"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
	"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
	"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
	"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
	"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
	"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
	"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
	"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
	"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
	"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
}

func init() {
	rand.Seed(time.Now().UnixNano())
	KEY = make([]byte, BS)
	rand.Read(KEY)
	IV = make([]byte, BS)
	rand.Read(IV)

	s = inputStrings[rand.Intn(len(inputStrings))]
}

func first(in string) []byte {
	// select at random one of the following 10 strings

	// generate a random AES key (which it should save for all future encryptions)

	// pad the string out to the 16-byte AES block size

	// CBC-encrypt it under that key
	// return the ciphertext and IV.

	return tools.EncryptAesCBC([]byte(in), KEY, IV)
}

func second(iv, ciphertext []byte) bool {
	// decrypt the ciphertext produced by the first function
	_, err := tools.DecryptAesCBC(ciphertext, KEY, iv)
	// check its padding
	if err != nil {
		return false
	}
	// return true or false depending on whether the padding is valid.
	return true
}

func findLastByte(prevBlock, targetBlock []byte) (byte, error) {
	return findByteN(prevBlock, targetBlock, 1)
	// copyPrevBlock := append([]byte{}, prevBlock...)
	// for c := 0; c < 256; c++ {
	// 	copyPrevBlock[len(copyPrevBlock)-1] = byte(c)
	// 	if second(copyPrevBlock, targetBlock) {
	// 		return 0x01 ^ byte(c) ^ prevBlock[len(prevBlock)-1], nil
	// 	}
	// }

	// return 0, errors.New("Error")
}

// ......1

// .....22

// ....3..

// .....22

func findByteN(prevBlock, targetBlock []byte, n int) (byte, error) {

	copyPrevBlock := append([]byte{}, prevBlock...)
	for c := 0; c < 256; c++ {
		copyPrevBlock[len(copyPrevBlock)-n] = byte(c)
		if second(copyPrevBlock, targetBlock) {
			return byte(n) ^ byte(c) ^ prevBlock[len(prevBlock)-n], nil
		}
	}

	return 0, errors.New("Error")
}

func findBlockBytes(prevBlock, targetBlock []byte) []byte {
	result := make([]byte, BS)
	copyPrevBlock := append([]byte{}, prevBlock...)
	for n := 1; n < BS+1; n++ {
		for m := n - 1; m > 0; m-- {
			copyPrevBlock[BS-m] = byte(n) ^ result[BS-m] ^ prevBlock[BS-m]
		}

		b, err := findByteN(copyPrevBlock, targetBlock, n)

		if err != nil {
			fmt.Println("error")
		} else {
			result[len(result)-n] = b
			// fmt.Printf("Got %s\n", string([]byte{b}))
		}
	}
	return result
}

func main() {
	s := inputStrings[rand.Intn(len(inputStrings))]
	fmt.Println("original", s)
	encrypted := first(s)
	fmt.Println("First block")
	result := findBlockBytes(IV, encrypted[0:BS])
	fmt.Printf("Should be: %s, got %s\n", s[0:BS], result)
	blockCount := int(math.Ceil(float64(len(encrypted)) / float64(BS)))
	for n := 1; n < blockCount-1; n++ {
		fmt.Println("Block #", n)
		result := findBlockBytes(encrypted[(n-1)*BS:n*BS], encrypted[n*BS:(n+1)*BS])
		fmt.Printf("Should be: %s, got %s\n", s[n*BS:(n+1)*BS], result)
	}

}
