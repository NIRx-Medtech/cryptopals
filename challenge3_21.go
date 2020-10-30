package main

import (
	"fmt"

	"github.com/robquant/cryptopals/pkg/tools"
)

func main() {
	fmt.Println("---------------")
	// https://pkg.go.dev/github.com/seehuhn/mt19937

	ourRandGen := tools.NewMT19937Rng(1000)
	ourVal := ourRandGen.Random()
	fmt.Println(ourVal)
	ourVal = ourRandGen.Random()
	fmt.Println(ourVal)
	fmt.Println("!###")
}

// https://godbolt.org/z/4vdTPd

// // maersenne_twister_engine::seed example
// #include <iostream>
// #include <random>

// int main ()
// {

//   std::mt19937 generator (1000);   // mt19937 is a standard mersenne_twister_engine
// for (int i = 0; i < 10; i++)
//   std::cout << generator() << std::endl;

//   return 0;
// }
