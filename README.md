# Bitcoin Miner

A simple, concurrent bitcoin miner framework implemented in Go.

> #### Disclaimer: this is not a product intended to be used for real mining, since I have little expertise on the topic. We do not take any responsibility or liability around any damage, financial loss or loss of any kind caused or in any way linked to this software. See the LICENSE file for more information.

## Installation

```shell
go get github.com/password-classified/bitcoinminer
```

Then, import it with...

```go
import (
	m "github.com/password-classified/bitcoinminer/miner"
)
```

<br>

## Usage

The function `miner.NewMiner` returns a miner struct that has the method `Mine`.

Pass the transaction string to the first argument of `NewMiner`, and the difficulty (needed number of 0s at the start) as the second argument.

```go
miner := m.NewMiner(`[BLOCK NUMBER]
[Transaction 1]
[Transaction 2]
[Transaction 3]
[Transaction ...]
`+m.SHA256("prev_block"), 6)
```

For mining, the syntax of the `Miner.Mine` method is as follows.

```go
miner.Mine(
	iterations_per_thread, // An integer
	threads, // An integer
	verbose_output, // A boolean
)
```
<br>

### Examples
To conclude, here is an example script...

```go
package main

import (
	m "github.com/password-classified/bitcoinminer/miner"
)

func main() {
	miner := m.NewMiner(`123
Bob->Steve->20
Gerald->Mary->14
Angela->Axel->120
`+m.SHA256("prev_block_hash_here"), 6)

	miner.Mine(500000, 3000, true)
}
```