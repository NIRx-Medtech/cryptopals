package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/robquant/cryptopals/pkg/tools"
)

// Write a routine that performs the following operation:
func generator() {
	// Wait a random number of seconds between, I don't know, 40 and 1000.
	// Seeds the RNG with the current Unix timestamp
	// Waits a random number of seconds again.
	// Returns the first 32 bit output of the RNG.
}

// You get the idea. Go get coffee while it runs. Or just simulate the passage of time, although you're missing some of the fun of this exercise if you do that.

// From the 32 bit RNG output, discover the seed.

func main() {
	fmt.Println("Ch3/22---------------")
	time.Sleep(time.Duration(5 + rand.Int31n(5)))
	timestamp := int(time.Now().Unix())
	ourRandGen := tools.NewMT19937Rng(timestamp)
	time.Sleep(time.Duration(5 + rand.Int31n(5)))

	ourVal := ourRandGen.Random()
	fmt.Println(ourVal)
	start := int(time.Now().Unix())

	for i := uint(0); i < ^uint(0); i++ {
		newSeed := start - int(i)
		rnd := tools.NewMT19937Rng(newSeed)
		if rnd.Random() == ourVal {
			fmt.Printf("Found seed: %d", newSeed)
			break
		}
	}
}
