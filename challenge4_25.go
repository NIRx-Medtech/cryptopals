package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/robquant/cryptopals/pkg/tools"
)

const BS = 16

func main() {
	key := []byte("YELLOW SUBMARINE")

	inputBytes, err := ioutil.ReadFile("input/input1_7.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Replace(string(inputBytes), "\n", "", 0)
	ciphertext, _ := base64.StdEncoding.DecodeString(input)
	plaintext := tools.DecryptAesECB(ciphertext, key)

	rand.Seed(time.Now().UnixNano())
	aesKey := make([]byte, BS)
	rand.Read(aesKey)
	nonce := make([]byte, 16)
	ctr := tools.NewCTR(nonce, aesKey)
	newciphertext := ctr.Encrypt(plaintext)

	work := append(make([]byte, 0), newciphertext...)
	ctr.Edit(work, 0, make([]byte, len(newciphertext)))
	newplaintext, _ := tools.Xor(newciphertext, work)
	fmt.Println(string(newplaintext))
	fmt.Println("==================")
	fmt.Println(string(plaintext))

}

//  1 2 3 4 5 6 7 8 9
//  2 AB
// 1 2 A' B' 5 6 7 8 9
