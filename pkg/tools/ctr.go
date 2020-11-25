package tools

import (
	"encoding/binary"
	"log"
	"math"
)

const BS = 16

type KeyStream interface {
	NextByte() byte
}

func CtrStream(input []byte, ks KeyStream) []byte {
	xored := make([]byte, len(input))
	for i, b := range input {
		xored[i] = b ^ ks.NextByte()
	}
	return xored
}

type CTR struct {
	counter int
	nonce   []byte
	xorKey  []byte
	aesKey  []byte
}

func NewCTR(nonce, aesKey []byte) *CTR {
	return &CTR{counter: 0, nonce: nonce, xorKey: nil, aesKey: aesKey}
}

func (ctr *CTR) nextXorKey() []byte {
	combined := make([]byte, 16)
	copy(combined, ctr.nonce)
	binary.LittleEndian.PutUint64(combined[8:], uint64(ctr.counter))
	//EncryptAES(plaintext, key []byte)
	xorKey := EncryptAES(combined, ctr.aesKey)
	ctr.counter++
	ctr.xorKey = append(ctr.xorKey, xorKey...)
	return xorKey
}

func (ctr *CTR) Reset() {
	ctr.counter = 0
	ctr.xorKey = make([]byte, 0)
}

func (ctr *CTR) Encrypt(plaintext []byte) []byte {
	blockCount := int(math.Ceil(float64(len(plaintext)) / float64(BS)))
	encryptedBytes := make([]byte, 0)
	for blockIndex := 0; blockIndex < blockCount-1; blockIndex++ {
		ctr.nextXorKey()
		xored, err := Xor(plaintext[blockIndex*BS:(blockIndex+1)*BS], ctr.xorKey[:BS])
		ctr.xorKey = ctr.xorKey[BS:]
		if err != nil {
			log.Fatal(err)
		}
		encryptedBytes = append(encryptedBytes, xored...)
	}
	//last block
	lenLastBlock := len(plaintext) % BS
	ctr.nextXorKey()
	xored, err := Xor(plaintext[len(plaintext)-lenLastBlock:], ctr.xorKey[:lenLastBlock])
	ctr.xorKey = ctr.xorKey[lenLastBlock:]
	if err != nil {
		log.Fatal(err)
	}
	encryptedBytes = append(encryptedBytes, xored...)
	return encryptedBytes
}

func (ctr *CTR) Edit(ciphertext []byte, offset int, newtext []byte) {
	ctr.Reset()
	garbage := make([]byte, offset)
	ctr.Encrypt(garbage)
	copy(ciphertext[offset:], ctr.Encrypt(newtext))
}
