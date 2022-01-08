package miner_test

import (
	"testing"

	b "github.com/password-classified/bitcoinminer/miner"
)

func TestMine(t *testing.T) {
	miner := b.NewMiner(`123
Bob->Steve->20
Gerald->Mary->14
Angela->Axel->120
`+b.SHA256("prev_block"), 6)

	miner.Mine(500000, 20, true)
}
