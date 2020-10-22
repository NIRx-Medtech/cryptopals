package main

import (
	"fmt"
)

const (
	w            = 32
	n            = 624
	m            = 397
	r            = 31
	a            = 0x9908B0DF
	u            = 11
	d            = 0xFFFFFFFF
	s            = 7
	b            = 0x9D2C5680
	t            = 15
	c            = 0xEFC60000
	l            = 18
	f            = 1812433253
	lowerMask    = (1 << r) - 1 // That is, the binary number of r 1's
	upperMask    = (^lowerMask) & ((1 << w) - 1)
	lowerBitsOfW = ((1 << w) - 1)
)

type MT19937Rng struct {
	MT    [n]int // state of the generator
	index int
}

func NewMT19937Rng(seed int) *MT19937Rng {
	var mt = MT19937Rng{
		index: n,
	}
	// Initialize the generator from a seed
	mt.MT[0] = seed
	for i := 1; i < n; i++ { // loop over each element
		mt.MT[i] = lowerBitsOfW & (f*(mt.MT[i-1]^(mt.MT[i-1]>>(w-2))) + i)
	}
	return &mt
}

func (mt *MT19937Rng) Random() int {
	if mt.index == n {
		mt.twist()
	}
	y := mt.MT[mt.index]
	y = y ^ ((y >> u) & d)
	y = y ^ ((y << s) & b)
	y = y ^ ((y << t) & c)
	y = y ^ (y >> l)

	mt.index = mt.index + 1
	return lowerBitsOfW & y
}

func (mt *MT19937Rng) twist() {
	var x int
	var xA int
	for i := 0; i < n; i++ {
		x = (mt.MT[i] & upperMask) + (mt.MT[(i+1)%n] & lowerMask)
		xA = x >> 1
		if (x % 2) != 0 { // lowest bit of x is 1
			xA = xA ^ a
		}
		mt.MT[i] = mt.MT[(i+m)%n] ^ xA
	}
	mt.index = 0
}

func main() {
	fmt.Println("---------------")
	// https://pkg.go.dev/github.com/seehuhn/mt19937

	ourRandGen := NewMT19937Rng(1000)
	ourVal := ourRandGen.Random()
	fmt.Println(ourVal)
	ourVal = ourRandGen.Random()
	fmt.Println(ourVal)
	fmt.Println("!###")
}
