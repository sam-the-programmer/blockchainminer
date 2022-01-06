package main

import (
	m "github.com/password-classified/bitcoinminer/miner"
)

func main() {
	miner := m.NewMiner(`123
Bob->Steve->20
Gerald->Mary->14
Angela->Axel->120
`+m.SHA256("prev_block"), 6)

	miner.Mine(500000, 3000, true, false)
}
