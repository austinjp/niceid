<hr>
<div align="center">

This is a port of [the original](https://github.com/RienNeVaPlus/human-id) to Go.

</div>
<hr>


## Nice Identifiers

NiceID generates readable strings by chaining common short English words in a semi-meaningful way.
The result is concatenated of `adjective + adjective + noun + verb + adverb` resulting in a pool size of over **50,850,000,000** possible combinations.

- **SFW**: no bad words; family friendly results
- No dependencies


## Examples

- some-loose-roses-crash-safely
- late-moody-stars-slide-brightly
- old-true-baths-relate-busily
- orange-floppy-lines-wink-freely


## Usage

```bash
go get github.com/austinjp/niceid
```

```go
package main

import (
	"fmt"
	"github.com/austinjp/niceid"
)

func main() {
	options := niceid.Options{}
	id := niceid.ID(options)
	fmt.Println(id)

	// Change the word delimeter (default is "-")
	options.Delim = " "
	fmt.Println(id)

	// For a deterministic outcome pass a seeded rand:
	deterministicOptions := Options{
		Rand: rand.New(rand.NewPCG(1, 2)),
	}
	deterministicID := niceid.ID(options)
	fmt.Println(deterministicID)
}

```
