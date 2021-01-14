package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/robquant/cryptopals/pkg/tools"
)

var key, nonce []byte
var prefix = []byte("comment1=cooking%20MCs;userdata=")
var postfix = []byte(";comment2=%20like%20a%20pound%20o%20bacon")

func init() {
	rand.Seed(time.Now().UnixNano())
	key = make([]byte, 16)
	rand.Read(key)
	nonce = make([]byte, 16)
	rand.Read(nonce)
}

func first(in string) []byte {
	combined := make([]byte, 0)
	// prepend the string: "comment1=cooking%20MCs;userdata="
	combined = append(combined, prefix...)
	// quote out the ";" and "=" characters.
	quoted := strings.ReplaceAll(in, ";", "\";\"")
	quoted = strings.ReplaceAll(quoted, "=", "\"=\"")
	// input string, quoted out, in the middle
	combined = append(combined, []byte(quoted)...)
	// append the string: ";comment2=%20like%20a%20pound%20of%20bacon"
	combined = append(combined, postfix...)

	ctr := tools.NewCTR(nonce, key)
	// encrypt it under the random AES key.
	return ctr.Encrypt(combined)
}

func second(ciphertext []byte) bool {
	// decrypt the string and look for the characters ";admin=true;"

	ctr := tools.NewCTR(nonce, key)
	// encrypt it under the random AES key.
	decryptedBytes := ctr.Encrypt(ciphertext)
	decrypted := string(decryptedBytes)

	// return true or false based on whether the string exists.
	return strings.Contains(decrypted, ";admin=true;")
}

func main() {
	encrypted1 := first("A")
	encrypted2 := first("B")
	var i int
	for i = 0; encrypted1[i] == encrypted2[i]; i++ {
	}
	prefixLength := i
	cleanPlaintext := "_admin_true_"
	wantedPlaintext := []byte(";admin=true;")
	encrypted := first(cleanPlaintext)
	keystreamBytes, _ := tools.Xor([]byte(cleanPlaintext), encrypted[prefixLength:prefixLength+len(cleanPlaintext)])
	corruptCipertext, _ := tools.Xor(keystreamBytes, wantedPlaintext)
	copy(encrypted[prefixLength:], corruptCipertext)
	fmt.Printf("%v", second(encrypted))
}
