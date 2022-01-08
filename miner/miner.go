package miner

import (
	sha "crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Miner struct {
	transactionString string
	Difficulty        int
}

// SHA256 returns the Sha256 of as string as a string
func SHA256(s string) string {
	return fmt.Sprintf("%x", sha.Sum256([]byte(s)))
}

// Mine returns the nonce value that works.
func (m *Miner) Mine(iterations int, threads int, hashTwice bool, output bool) int {
	var hashed string

	solutions := make(chan []int, 100)

	fmt.Println("Preparing to test", threads*iterations, "nonce values to find a difficulty of", m.Difficulty)

	// For Concurrency
	for t := 0; t < threads; t++ {
		if output {
			fmt.Println("Thread", t, "\b: starting search at", t*iterations, "until", t*iterations+iterations)
		}

		// Mining
		go func(start int, c chan<- []int, t int) {

			nonce := start

			for i := 0; i < iterations; i++ {
				if hashTwice {
					hashed = SHA256(SHA256(m.transactionString + strconv.Itoa(nonce+i)))
				} else {
					hashed = SHA256(m.transactionString + strconv.Itoa(nonce+i))
				}

				if hashed[:m.Difficulty] == strings.Repeat("0", m.Difficulty) {
					c <- []int{nonce + i, t}
					break
				}
			}
		}(t*iterations, solutions, t) // So that it can scan many integers simultaneously, but still not miss any
	}

	fmt.Println("All", threads, "threads have been created. Searching for values.")

	n := <-solutions
	fmt.Println("Successful trial in thread", n[1], "with nonce value =", n[0])

	return n[0]
}

func NewMiner(transaction string, difficulty int) Miner {
	return Miner{
		transactionString: transaction,
		Difficulty:        difficulty,
	}
}
