package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/robquant/cryptopals/pkg/tools"
)

type MT19937RngKeyStream struct {
	rng     tools.MT19937Rng
	current uint32
	index   int
}

func NewMT19937RngKeyStream(seed uint16) *MT19937RngKeyStream {
	return &MT19937RngKeyStream{rng: *tools.NewMT19937Rng(int(seed)), index: -1}
}

func (ks *MT19937RngKeyStream) NextByte() byte {
	if ks.index == -1 || ks.index == 4 {
		ks.current = uint32(ks.rng.Random())
		ks.index = 0
	}
	res := byte(ks.current >> (ks.index * 8))
	ks.index++
	return res
}

func main() {
	fmt.Println("Ch3/24---------------")
	rand.Seed(time.Now().UnixNano())
	key := uint16(rand.Uint32())
	ks := NewMT19937RngKeyStream(key)
	prefixLength := 5 + rand.Int31n(20)
	message := make([]byte, prefixLength)
	rand.Read(message)
	message = append(message, []byte("AAAAAAAAAAAAAA")...)
	ciphertext := tools.CtrStream(message, ks)
	for guess := uint16(0); guess <= 65535; guess++ {
		ks := NewMT19937RngKeyStream(guess)
		possiblePlaintext := tools.CtrStream(ciphertext, ks)
		if strings.HasSuffix(string(possiblePlaintext), "AAAAAAAAAAAAAA") {
			fmt.Printf("Found key: %d\n", guess)
			break
		}
	}
	fmt.Printf("The real key was: %d\n", key)
}
