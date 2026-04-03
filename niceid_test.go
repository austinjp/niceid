package niceid

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func TestPrintID(t *testing.T) {
	var opts = Options{}
	fmt.Println(ID(opts))
	fmt.Println(ID(opts))
	fmt.Println(ID(opts))
}

func TestPrintDeterministicID(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	var opts = Options{Rand: r}
	fmt.Println(ID(opts))
	fmt.Println(ID(opts))

	opts.Delim = " "
	fmt.Println(ID(opts))
	fmt.Println(ID(opts))
}

func TestPrintNumberOfIDs(t *testing.T) {
	fmt.Println(len(Adjectives) * (len(Adjectives) - 11) * len(Nouns) * len(Verbs) * len(Adverbs))
}
