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

func InvertRightShiftOp(y, l int32) {
	res := int32(0)

	// 1010
	// 0101
	// 1111

}

func main() {
	fmt.Println("Ch3/23---------------")
	// y = y ^ ((y >> u) & d)
	// y = y ^ ((y << s) & b)
	// y = y ^ ((y << t) & c)
	// y = y ^ (y >> l)
	// return lowerBitsOfW & y
}
