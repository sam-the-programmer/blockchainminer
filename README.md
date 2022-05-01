# Bitcoin Miner

A simple, concurrent bitcoin miner framework implemented in Go.

> #### Disclaimer: this is not a product intended to be used for real mining, since I have little expertise on the topic. We do not take any responsibility or liability around any damage, financial loss or loss of any kind caused or in any way linked to this software. See the LICENSE file for more information.

## Installation

```shell
go get github.com/sam-the-programmer/bitcoinminer
```

Then, import it with...

```go
import (
	m "github.com/sam-the-programmer/bitcoinminer/miner"
)
```

<br>

## Usage

The bulk of the API is here. More detail is on [pkg.go.dev](https://pkg.go.dev/github.com/sam-the-programmer/bitcoinminer)

```go
package main

import (
	"fmt"

	"github.com/sam-the-programmer/bitcoinminer/hash"
	"github.com/sam-the-programmer/bitcoinminer/miner"
)

func main() {
	transactionString := "<BlockNum-0>\nAlice->Bob->20\nBob->Charlie->10\nCharlie->Alice->5\n[PrevBlockHash]\n%v"
	m := miner.NewMiner(
		transactionString,
		hash.SHA256,
		10000,
	)

	m.SetSearchSize(100000000000)
	m.SetDifficulty(6)
	m.SetHashTimes(1)
	m.SetOutputLevel(1)

	solution, hasFound := m.ThreadedMine()
	if hasFound {
		fmt.Println("\n\n", solution, "\n", hash.SHA256(hash.SHA256(fmt.Sprintf(transactionString, solution))))
	}
}

```

<br>

