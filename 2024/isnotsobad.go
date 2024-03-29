package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for range 10 {
		fmt.Println("give me a rand number! >", rand.N(int8(120)))
	}
}
