package main

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var one = big.NewInt(1)

func pow(num, p int64) *big.Int {
	var result big.Int
	return result.Exp(big.NewInt(num), big.NewInt(p), nil)
}

func generateRandomKey(n int64) *big.Int {
	random := big.NewInt(0)
	for i := int64(0); i < n; i++ {
		randBit := rand.Intn(2)
		if randBit > 0 {
			random.Add(random, pow(int64(2), i))
		}
	}
	return random
}

func bruteForceKey(target *big.Int, n int64) (time.Duration, error) {
	numOfKeys := pow(2, n)
	startTime := time.Now()
	for i := big.NewInt(0); i.Cmp(numOfKeys) < 0; i.Add(i, one) {
		if i.Cmp(target) == 0 {
			return time.Since(startTime), nil
		}
	}
	return 0, errors.New("match not found")
}

func main() {
	const size = 10
	rand.Seed(time.Now().UnixNano())

	bits := [size]int64{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
	var randKeys [size]*big.Int

	fmt.Println("[1] Calculate number of n-bit keys")
	for i := 0; i < len(bits); i++ {
		numOfKeys := pow(2, bits[i])
		fmt.Printf("Number of keys for n=%d: %s\n", bits[i], numOfKeys)
	}

	fmt.Println("\n[2] Generate a random key from an n-bits keys field")
	for i := 0; i < len(bits); i++ {
		randKey := generateRandomKey(bits[i])
		randKeys[i] = randKey
		fmt.Printf("Random key from %d-bits keys field: 0x%x\n", bits[i], randKey)
	}

	fmt.Println("\n[3] Bruteforce keys until the random key from [2] is found")
	for i := 0; i < len(bits); i++ {
		timePassed, err := bruteForceKey(randKeys[i], bits[i])
		if err != nil {
			fmt.Printf("Match for n=%d not found\n", bits[i])
		} else {
			fmt.Printf("Brute force algorithm for n=%d found a key in %d ms\n",
				bits[i], timePassed.Milliseconds())
		}
	}
}
