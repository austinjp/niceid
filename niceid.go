package niceid

import (
	"math/rand/v2"
	"strings"
)

// Options must be passed to ID. Fields:
//
//	Delim: word delimeter (default "-")
//	Rand:  pass *rand.Rand for deterministic output, otherwise
//		math/rand/v2 is used
type Options struct {
	Delim string
	Rand  *rand.Rand
}

// ID generates a human-friendly identifier of the form:
//
//	<adjective>-<adjective>-<noun>-<verb>-<adverb>
func ID(o Options) string {
	var intN func(n int) int
	if o.Rand == nil {
		intN = rand.IntN
	} else {
		intN = rand.New(o.Rand).IntN
	}
	var delim string
	if o.Delim == "" {
		delim = "-"
	} else {
		delim = o.Delim
	}

	lenAdjectives := len(Adjectives)

	a1 := Adjectives[intN(lenAdjectives)]
	var a2 string

	// Make sure a1 and a2 are different and a2 isn't a quantity.
	avoid := map[string]bool{
		a1:      true,
		"some":  true,
		"two":   true,
		"three": true,
		"four":  true,
		"five":  true,
		"six":   true,
		"seven": true,
		"eight": true,
		"nine":  true,
		"ten":   true,
	}
	for {
		a2 = Adjectives[intN(lenAdjectives)]
		if len(avoid) >= lenAdjectives {
			break
		}
		if avoid[a2] {
			continue
		}
		break
	}

	n := Nouns[intN(len(Nouns))]
	v := Verbs[intN(len(Verbs))]
	d := Adverbs[intN(len(Adverbs))]

	var b strings.Builder
	b.WriteString(a1)
	b.WriteString(delim)
	b.WriteString(a2)
	b.WriteString(delim)
	b.WriteString(n)
	b.WriteString(delim)
	b.WriteString(v)
	b.WriteString(delim)
	b.WriteString(d)
	return b.String()
}
