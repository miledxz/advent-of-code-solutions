package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/miledxz/advent-of-code-solutions/utils"
)

func Solve1() any {
	lns := utils.ReadLines("07.txt")

	var g G
	for _, ln := range lns {
		h := NewH(ln)
		g = append(g, h)
	}

	sort.Sort(g)

	var sum int64
	for i, v := range g {
		sum += int64(i+1) * v.Bid
	}
	return sum
}

func Solve2() any {
	lns := utils.ReadLines("07.txt")

	var g G
	for _, ln := range lns {
		h := NewH(ln)
		g = append(g, h)
	}

	sort.Sort(g)

	var sum int64
	for i, v := range g {
		sum += int64(i+1) * v.Bid
	}
	return sum
}

type H int64

const (
	HighCard H = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Type  H
	Cards []rune
	Bid   int64
}

func NewH1(ln string) *Hand {
	parts := strings.Fields(ln)

	cards := []rune(parts[0])
	bid, err := strconv.ParseInt(parts[1], 10, 64)
	utils.Check(err, "Unable to parse %s to int64", parts[1])

	counts := make(map[rune]int)

	for _, c := range cards {
		counts[c]++
	}

	cardcounts := make([]int, 0, 5)

	for _, v := range counts {
		cardcounts = append(cardcounts, v)
	}

	sort.Ints(cardcounts)

	var h H
	switch {
	case equal(cardcounts, []int{5}):
		h = FiveOfAKind
	case equal(cardcounts, []int{1, 4}):
		h = FourOfAKind
	case equal(cardcounts, []int{2, 3}):
		h = FullHouse
	case equal(cardcounts, []int{1, 1, 3}):
		h = ThreeOfAKind
	case equal(cardcounts, []int{1, 2, 2}):
		h = TwoPair
	case equal(cardcounts, []int{1, 1, 1, 2}):
		h = OnePair
	case equal(cardcounts, []int{1, 1, 1, 1, 1}):
		h = HighCard
	default:
		panic(fmt.Sprintf("Unexpected card counts: %v", cardcounts))
	}

	return &Hand{h, cards, bid}
}

func NewH(ln string) *Hand {
	parts := strings.Fields(ln)

	cards := []rune(parts[0])
	bid, err := strconv.ParseInt(parts[1], 10, 64)
	utils.Check(err, "Unable to parse %s to int64", parts[1])

	h := getBestHand(cards)

	return &Hand{h, cards, bid}
}

func getBestHand(cards []rune) H {
	counts := make(map[rune]int)

	for _, c := range cards {
		counts[c]++
	}

	if counts['J'] != 0 {
		maxv := 0
		var maxc rune
		for c, v := range counts {
			if c != 'J' && v > maxv {
				maxc = c
				maxv = v
			}
		}

		counts[maxc] += counts['J']
		counts['J'] = 0
	}

	cardcounts := make([]int, 0, 5)

	for _, v := range counts {
		if v != 0 {
			cardcounts = append(cardcounts, v)
		}
	}

	sort.Ints(cardcounts)

	var h H
	switch {
	case equal(cardcounts, []int{5}):
		h = FiveOfAKind
	case equal(cardcounts, []int{1, 4}):
		h = FourOfAKind
	case equal(cardcounts, []int{2, 3}):
		h = FullHouse
	case equal(cardcounts, []int{1, 1, 3}):
		h = ThreeOfAKind
	case equal(cardcounts, []int{1, 2, 2}):
		h = TwoPair
	case equal(cardcounts, []int{1, 1, 1, 2}):
		h = OnePair
	case equal(cardcounts, []int{1, 1, 1, 1, 1}):
		h = HighCard
	default:
		panic(fmt.Sprintf("Unexpected card counts: %v", cardcounts))
	}

	return h
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

type G [](*Hand)

var cardValues = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': -1,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func (g G) Len() int      { return len(g) }
func (g G) Swap(i, j int) { g[i], g[j] = g[j], g[i] }
func (g G) Less(i, j int) bool {
	a := g[i]
	b := g[j]

	if a.Type != b.Type {
		return a.Type < b.Type
	}

	for k := 0; k < 5; k++ {
		if cardValues[a.Cards[k]] != cardValues[b.Cards[k]] {
			return cardValues[a.Cards[k]] < cardValues[b.Cards[k]]
		}
	}

	return false
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)

}
