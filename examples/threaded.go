package main

import (
	"fmt"

	"github.com/sam-the-programmer/bitcoinminer/hash"
	"github.com/sam-the-programmer/bitcoinminer/miner"
)

const transactionString = "<BlockNum-0>\nAlice->Bob->20\nBob->Charlie->10\nCharlie->Alice->5\n[PrevBlockHash]\n%v"

func main() {

	m := miner.NewMiner(
		transactionString,
		hash.SHA256,
	)

	m.SetSearchSize(1000000000000)
	m.SetDifficulty(6)
	m.SetHashTimes(2)
	m.SetOutputLevel(1)

	solution := m.MineForever()
	// if hasFound {
	fmt.Println("\n\n", solution, "\n", m.MultiHashFunc(fmt.Sprintf(transactionString, solution)))
	// }
}
