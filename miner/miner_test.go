package miner_test

import (
	"reflect"
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

func TestMiner_SetHashTimes(t *testing.T) {
	type args struct {
		t uint
	}
	tests := []struct {
		name string
		m    *b.Miner
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetHashTimes(tt.args.t)
		})
	}
}

func TestMiner_Mine(t *testing.T) {
	type args struct {
		iterations int
		threads    int
		output     bool
	}
	tests := []struct {
		name string
		m    *b.Miner
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Mine(tt.args.iterations, tt.args.threads, tt.args.output); got != tt.want {
				t.Errorf("Miner.Mine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMiner(t *testing.T) {
	type args struct {
		transaction string
		difficulty  int
		hash        b.HashFunction
	}
	tests := []struct {
		name string
		args args
		want b.Miner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.NewMiner(tt.args.transaction, tt.args.difficulty, tt.args.hash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMiner() = %v, want %v", got, tt.want)
			}
		})
	}
}
