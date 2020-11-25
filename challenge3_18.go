package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/robquant/cryptopals/pkg/tools"
)

const BS = 16
const KEY = "YELLOW SUBMARINE"
const input = "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ=="

var nonce []byte

var counter int

func init() {
	counter = 0
	nonce = make([]byte, 8)
}

func main() {
	decodedInput, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		log.Fatal("Decoding error ", err)
	}

	ctr := tools.NewCTR(nonce, []byte(KEY))
	res := ctr.Encrypt(decodedInput[:10])
	fmt.Println(string(res))
	res = ctr.Encrypt(decodedInput[10:])
	fmt.Println(string(res))
	fmt.Println("---------------")
	ctr.Reset()
	res = ctr.Encrypt(decodedInput)
	fmt.Println(string(res))
}
