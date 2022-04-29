package miner_test

import (
	"testing"

	h "github.com/sam-the-programmer/bitcoinminer/hash"
	b "github.com/sam-the-programmer/bitcoinminer/miner"
)

func TestMine(t *testing.T) {
	miner := b.NewMiner(`123
Bob->Steve->20
Gerald->Mary->14
Angela->Axel->120
`+h.SHA256("prev_block"), 6, h.SHA256)

	miner.Mine(500000, 20, true)
}
