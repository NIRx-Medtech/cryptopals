package main

import (
	"fmt"

	"github.com/robquant/cryptopals/pkg/tools"
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

func IsolateBit(val uint32, bit uint8) uint32 {
	return val & (1 << bit)
}

// Revert y = y ^ ((y >> shift) & a)
func RevertRightShiftOp(y uint32, shift uint8, a uint32) uint32 {
	res := uint32(0)
	for i := uint8(0); i < shift; i++ {
		res |= IsolateBit(y, 31-i)
	}
	for i := shift; i < 32; i++ {
		unshifted := IsolateBit(res, 31-i+shift)
		res |= IsolateBit(y, 31-i) ^ ((unshifted >> shift) & a)
	}
	return res
}

// Revert y = y ^ ((y << shift) & a)
func RevertLeftShiftOp(y uint32, shift uint8, a uint32) uint32 {
	res := uint32(0)
	for i := uint8(0); i < shift; i++ {
		res |= IsolateBit(y, i)
	}
	for i := shift; i < 32; i++ {
		unshifted := IsolateBit(res, i-shift)
		res |= IsolateBit(y, i) ^ ((unshifted << shift) & a)
	}
	return res
}

//shift = 2
//	x2 x3 x4 x5 0  0
// &
//	a0 a1 a2 a3 a4 a5
// ^
//	x0 x1 x2 x3 x4 x5
// -------------------
//  y0 y1 y2 y3 y4 y5

//res[i] = y[i]
//res[i] = y[i] ^ (res[i+shift] & a[i])

// Revert y = y ^ ((y << shift) & a)
func RevertLeftShiftOp2(y uint32, shift uint8, a uint32) uint32 {
	res := uint32(0)
	numOfChunks := 32 / shift

	//mask with ones(length = shift) in the last position  000011
	baseMask := uint32(0xffffffff >> (32 - shift))

	for i := uint8(0); i <= numOfChunks; i++ {
		//mask shifted to the next bits of interest  i = 0 : 000011 ; i = 1 : 001100
		maskForChunk := uint32(baseMask << (shift * i))

		//new portion of final result
		chunk := y & maskForChunk // y4 y5 for i == 0

		//y[i] ^ (res[i+shift] & a[i])
		y ^= (chunk << shift) & a // y4 y5 (= x4 x5) shifted left for i == 0

		res |= chunk
	}
	return res
}

// y := mt.MT[mt.index]
// y = y ^ ((y >> u) & d)
// y = y ^ ((y << s) & b)
// y = y ^ ((y << t) & c)
// y = y ^ (y >> l)

func main() {
	rng := tools.NewMT19937Rng(100)
	rng2 := tools.NewMT19937Rng(138400)
	for i := 0; i < 624; i++ {
		rng2.Random()
		y := uint32(rng.Random())
		y = RevertRightShiftOp(y, l, ^uint32(0))
		y = RevertLeftShiftOp2(y, t, c)
		y = RevertLeftShiftOp2(y, s, b)
		y = RevertRightShiftOp(y, u, d)
		rng2.MT[i] = int(y)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Orig: %d, Copy: %d\n", rng.Random(), rng2.Random())
	}
}
