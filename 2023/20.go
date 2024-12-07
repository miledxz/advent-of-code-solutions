package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/miledxz/advent-of-code-solutions/utils"
)

func Solve1() any {
	lns := utils.ReadLines("20.txt")

	mods := make(map[string]*mod)
	for _, ln := range lns {
		n, mod := prsln(ln)
		mods[n] = mod
	}

	for n, mod := range mods {
		for _, d := range mod.destinations {
			if md, ok := mods[d]; ok {
				md.sources = append(md.sources, n)
			}
		}
	}

	bitpos := 0
	for _, mod := range mods {
		switch mod.kind {
		case ud:
			mod.bnum = bitpos
			bitpos++
		case conjunction:
			mod.bnum = bitpos
			bitpos += len(mod.sources)
		}
	}

	bits := NewB(bitpos)
	var lowC, highC uint64

	sb := bits.Serialize()
	for i := 0; i < 1000; i++ {
		res := pressB(sb, mods)
		sb = res.serializedBits
		lowC += res.lowPulses
		highC += res.highPulses
	}

	return lowC * highC
}

type pressResult struct {
	serializedBits string
	lowPulses      uint64
	highPulses     uint64
}

type pulse struct {
	source      string
	destination string
	isHigh      bool
}

func pressB(serializedState string, mods map[string]*mod) pressResult {
	bits, err := Deserialize(serializedState)
	utils.Check(err, "deserialization error for %s", serializedState)

	var lowC uint64
	var highC uint64

	queue := []pulse{{"button", "broadcaster", false}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if utils.Verbose {
			fmt.Printf("Pulse: %s -> %s  high: %t\n", p.source, p.destination, p.isHigh)
		}

		switch p.isHigh {
		case true:
			highC++
		case false:
			lowC++
		}

		md, ok := mods[p.destination]

		if !ok {
			continue
		}

		switch md.kind {
		case broadcast:
			for _, dn := range md.destinations {
				queue = append(queue, pulse{p.destination, dn, p.isHigh})
			}

		case ud:
			if p.isHigh == false {
				currentState := bits.Get(md.bnum)
				newState := !currentState
				switch newState {
				case true:
					bits.Set(md.bnum)
				case false:
					bits.Unset(md.bnum)
				}

				for _, dn := range md.destinations {
					queue = append(queue, pulse{p.destination, dn, newState})
				}
			}
		case conjunction:
			idx := -1
			for i, v := range md.sources {
				if v == p.source {
					idx = i
					break
				}
			}
			if idx < 0 {
				panic(fmt.Errorf("did not find %s in sources for mod %s", p.source, p.destination))
			}

			switch p.isHigh {
			case true:
				bits.Set(md.bnum + idx)
			case false:
				bits.Unset(md.bnum + idx)
			}

			sendHigh := false
			for i := md.bnum; i < md.bnum+len(md.sources); i++ {
				if !bits.Get(i) {
					sendHigh = true
					break
				}
			}
			for _, dn := range md.destinations {
				queue = append(queue, pulse{p.destination, dn, sendHigh})
			}
		}
	}

	res := pressResult{bits.Serialize(), lowC, highC}
	return res
}

type modtype int

const (
	ud modtype = iota
	conjunction
	broadcast
)

type mod struct {
	kind         modtype
	destinations []string
	sources      []string
	bnum         int
}

func prsln(ln string) (string, *mod) {
	var ret mod
	var n string

	components := strings.Split(ln, " -> ")
	ret.destinations = strings.Split(components[1], ", ")

	switch components[0][0] {
	case '%':
		ret.kind = ud
		n = components[0][1:]
	case '&':
		ret.kind = conjunction
		n = components[0][1:]
	default:
		ret.kind = broadcast
		n = components[0]
	}

	return n, &ret
}

type Bitfield struct {
	bits []uint64
}

func NewB(numBits int) *Bitfield {
	numUint64s := (numBits + 63) / 64
	return &Bitfield{
		bits: make([]uint64, numUint64s),
	}
}

func (bf *Bitfield) Set(bitIndex int) {
	wordIndex, offset := bitIndex/64, uint(bitIndex%64)
	bf.bits[wordIndex] |= (1 << offset)
}

func (bf *Bitfield) Get(bitIndex int) bool {
	wordIndex, offset := bitIndex/64, uint(bitIndex%64)
	return (bf.bits[wordIndex] & (1 << offset)) != 0
}

func (bf *Bitfield) Len() int {
	return len(bf.bits) * 64
}

func (bf *Bitfield) Unset(bitIndex int) {
	wordIndex, offset := bitIndex/64, uint(bitIndex%64)
	bf.bits[wordIndex] &^= (1 << offset)
}

func (bf *Bitfield) Serialize() string {
	ret := make([]string, len(bf.bits))

	for i, b := range bf.bits {
		ret[i] = strconv.FormatUint(b, 16)
	}

	return strings.Join(ret, ",")
}

func Deserialize(s string) (*Bitfield, error) {
	components := strings.Split(s, ",")

	bits := make([]uint64, len(components))

	for i, c := range components {
		v, err := strconv.ParseUint(c, 16, 64)
		if err != nil {
			return nil, err
		}

		bits[i] = v
	}
	return &Bitfield{bits}, nil
}

func Solve2() any {
	lns := utils.ReadLines("20.txt")

	mods := make(map[string]*mod)
	for _, ln := range lns {
		n, mod := prsln(ln)
		mods[n] = mod
	}

	for n, mod := range mods {
		for _, d := range mod.destinations {
			if md, ok := mods[d]; ok {
				md.sources = append(md.sources, n)
			}
		}
	}

	f, err := os.Create("day20p2/relations.gv")
	utils.Check(err, "unable to open visualization file")
	defer f.Close()

	fmt.Fprintln(f, "digraph G {")

	for n, m := range mods {
		var shape string
		switch m.kind {
		case ud:
			shape = "rect"
		case conjunction:
			shape = "trapezium"
		default:
			shape = "circle"
		}

		fmt.Fprintf(f, "  %s [ shape = \"%s\"; ];\n", n, shape)
		for _, d := range m.destinations {
			fmt.Fprintf(f, "  %s -> %s;\n", n, d)
		}
	}
	fmt.Fprintln(f, "}")

	val := int64(0b111101011011)
	val = utils.Lcm(val, int64(0b111100010111))
	val = utils.Lcm(val, int64(0b111011010101))
	val = utils.Lcm(val, int64(0b111010111001))

	return val
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)

}
