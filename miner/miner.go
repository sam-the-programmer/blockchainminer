// CPU coin miners.
package miner

import (
	"fmt"
	"strings"

	h "github.com/sam-the-programmer/bitcoinminer/hash"
)

// CPUMiner is a struct of a CPU coin miner.
type CPUMiner struct {
	Difficulty  uint   // The difficulty for the miner. Defaults to 1.
	HashTimes   uint   // The number of times to hash the value. Defaults to 1.
	Workers     uint64 // The number of threads to use for mining. Defaults to 1.
	SearchSize  uint64 // The number of nonces to search through. Defaults to 1000000000.
	OutputLevel uint8  // The level of output to give. Defaults to 0. 0 means nothing, 1 or above incrementally increases it. Do not set manually, use th Miner.SetOutputLevel method.

	Transaction   string         // The transaction to go through the hashing process.
	HashFunc      h.HashFunction // The hash function to use.
	MultiHashFunc h.HashFunction // The function to use for hashing multiple times.

	solutionChan    chan uint64 // The channel to send the nonce to.
	hasNotFoundChan chan bool   // The channel to tell if it has not been solved.

	solutionString string // The string to compare the hash to.

	threadOutputFunc   func(uint64, uint64)
	solutionOutputFunc func(uint64, bool)
}

// SetDifficulty sets the difficulty for the miner.
func (m *CPUMiner) SetDifficulty(t uint) {
	m.Difficulty = t
	m.solutionString = strings.Repeat("0", int(t))
}

// SetOutputLevel sets the level of output to give during mining.
// 0 means nothing,
// 1 or above incrementally increases output (max 1).
func (m *CPUMiner) SetOutputLevel(l int) {
	m.OutputLevel = uint8(l)
	switch l {
	case 0:
		m.threadOutputFunc = func(uint64, uint64) {}
		m.solutionOutputFunc = func(uint64, bool) {}
	default:
		m.threadOutputFunc = func(a uint64, b uint64) { fmt.Println("Thread\t", a+1, " searching\t", b, "\tvalues.") }
		m.solutionOutputFunc = func(value uint64, found bool) {
			if found {
				fmt.Println("Found solution value of ", value, ".")
			} else {
				fmt.Println("No nonce found after ", value, " iterations.")
			}
		}
	}
}

// SetSearchSize sets the number of nonces to search through.
func (m *CPUMiner) SetSearchSize(size uint64) {
	m.SearchSize = size
}

// SetWorkers sets the number of threads to use for mining.
func (m *CPUMiner) SetWorkers(workers uint64) {
	m.Workers = workers
	m.solutionChan = make(chan uint64, workers)
	m.hasNotFoundChan = make(chan bool, workers)
}

// SetMultiHashFunc sets the number of times to hash the value.
func (m *CPUMiner) SetHashTimes(t uint) {
	m.HashTimes = t
	m.MultiHashFunc = h.MultiHash(m.HashFunc, t)
}

// SetHashFunc sets hash function for mining.
func (m *CPUMiner) SetHash(f h.HashFunction) {
	m.HashFunc = f
	m.SetHashTimes(m.HashTimes)
}

// MineForever returns the nonce value for a mined block,
// but will search forever single-threadedly, never stopping.
func (m *CPUMiner) MineForever() uint64 {
	nonce := uint64(0)
	for {
		hash := m.MultiHashFunc(fmt.Sprintf(m.Transaction, nonce))
		if m.isValidHash(hash) {
			break
		}
		nonce++
	}

	m.solutionOutputFunc(nonce, true)
	return nonce
}

// ThreadedMine mines the block with the given difficulty,
// using multiple threads.
func (m *CPUMiner) ThreadedMine() (uint64, bool) {
	noncesPerWorker := uint64(m.SearchSize / uint64(m.Workers))
	for i := uint64(0); i < m.Workers; i++ {
		go m.mineThread(i*noncesPerWorker, noncesPerWorker)
		m.threadOutputFunc(i, noncesPerWorker)
	}

	workersDone := uint64(0)
	for {
		select {
		case nonce := <-m.solutionChan:
			m.solutionOutputFunc(nonce, true)
			return nonce, true
		case <-m.hasNotFoundChan:
			workersDone++
			if workersDone >= m.Workers {
				m.solutionOutputFunc(workersDone*noncesPerWorker, false)
				return 0, false
			}
		}
	}
}

// isValidHash returns true if the hash is valid given the
// miner's difficulty.
func (m *CPUMiner) isValidHash(hash string) bool {
	return m.solutionString == hash[:m.Difficulty]
}

// mineThread mines the block with the given difficulty,
// for the number of nonces stated. It can be used as a
// goroutine.
func (m *CPUMiner) mineThread(start uint64, num uint64) {
	for i := start; i < num; i++ {
		hash := m.MultiHashFunc(fmt.Sprintf(m.Transaction, i))
		if m.isValidHash(hash) {
			m.solutionChan <- i
		}
	}

	m.hasNotFoundChan <- true
}

// NewMiner creates a new miner struct, calling initialisation code.
func NewMiner(transaction string, hash h.HashFunction) CPUMiner {
	workers := uint64(100)

	return CPUMiner{
		Difficulty:  1,
		HashTimes:   1,
		Workers:     workers,
		SearchSize:  1000000000,
		OutputLevel: 0,

		Transaction:   transaction,
		HashFunc:      hash,
		MultiHashFunc: func(s string) string { return hash(s) },

		solutionChan:    make(chan uint64, workers),
		hasNotFoundChan: make(chan bool, workers),

		solutionString: strings.Repeat("0", 1),
	}
}
