package miner

import (
	"fmt"
	"strconv"
	"strings"
)

// HashFunction is a function that returns a hashed version of a string
type HashFunction func(string) string

// Miner is a struct of the mining code.
type Miner struct {
	Difficulty int
	HashTimes  uint

	transaction   string
	hashFunc      func(string) string
	multiHashFunc func(string) string
}

func (m *Miner) SetHashTimes(t uint) {
	m.HashTimes = t

	m.multiHashFunc = func(s string) string {
		hash := m.hashFunc(s)
		for i := uint(0); i < m.HashTimes-1; i++ {
			hash = m.hashFunc(hash)
		}
		return hash
	}
}

// Mine returns the nonce value that works.
func (m *Miner) Mine(iterations int, threads int, output bool) int {
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
				hashed = m.multiHashFunc(m.transaction + strconv.Itoa(nonce+i))

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

func NewMiner(transaction string, difficulty int, hash HashFunction) Miner {
	return Miner{
		transaction: transaction,
		Difficulty:  difficulty,
		hashFunc:    hash,
	}
}
