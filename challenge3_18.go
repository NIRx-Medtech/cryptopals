package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"math"

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

type CTR struct {
	counter int
	nonce   []byte
	xorKey  []byte
}

func (ctr *CTR) nextXorKey() []byte {
	combined := make([]byte, 16)
	copy(combined, ctr.nonce)
	binary.LittleEndian.PutUint64(combined[8:], uint64(ctr.counter))
	//EncryptAES(plaintext, key []byte)
	xorKey := tools.EncryptAES(combined, []byte(KEY))
	ctr.counter++
	ctr.xorKey = append(ctr.xorKey, xorKey...)
	return xorKey
}

func (ctr *CTR) reset() {
	ctr.counter = 0
	ctr.xorKey = make([]byte, 0)
}

func (ctr *CTR) Encrypt(plaintext []byte) []byte {
	blockCount := int(math.Ceil(float64(len(plaintext)) / float64(BS)))
	encryptedBytes := make([]byte, 0)
	for blockIndex := 0; blockIndex < blockCount-1; blockIndex++ {
		ctr.nextXorKey()
		xored, err := tools.Xor(plaintext[blockIndex*BS:(blockIndex+1)*BS], ctr.xorKey[:BS])
		ctr.xorKey = ctr.xorKey[BS:]
		if err != nil {
			log.Fatal(err)
		}
		encryptedBytes = append(encryptedBytes, xored...)
	}
	//last block
	lenLastBlock := len(plaintext) % BS
	ctr.nextXorKey()
	xored, err := tools.Xor(plaintext[len(plaintext)-lenLastBlock:], ctr.xorKey[:lenLastBlock])
	ctr.xorKey = ctr.xorKey[lenLastBlock:]
	if err != nil {
		log.Fatal(err)
	}
	encryptedBytes = append(encryptedBytes, xored...)
	return encryptedBytes
}

func main() {
	decodedInput, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		log.Fatal("Decoding error ", err)
	}

	ctr := &CTR{counter: 0, nonce: nonce, xorKey: nil}
	res := ctr.Encrypt(decodedInput[:10])
	fmt.Println(string(res))
	res = ctr.Encrypt(decodedInput[10:])
	fmt.Println(string(res))
	fmt.Println("---------------")
	ctr.reset()
	res = ctr.Encrypt(decodedInput)
	fmt.Println(string(res))
}
