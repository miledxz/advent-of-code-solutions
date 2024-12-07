package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/miledxz/advent-of-code-solutions/utils"
)

func Solve1() any {
	lns := utils.ReadLines("04.txt")

	ans := 0

	for _, ln := range lns {
		card := loadS(ln)

		w := make(map[int]bool)

		for _, v := range card.W {
			w[v] = true
		}

		res := 0

		for _, v := range card.S {
			if w[v] {
				res++
			}
		}

		switch res {
		case 0:
		default:
			ans += 1 << (res - 1)
		}
	}
	return ans
}

type Card struct {
	W []int
	S []int
}

func loadS(ln string) Card {
	fs := strings.Split(ln, ": ")
	c := strings.Split(fs[1], " | ")
	return Card{stringToNumbers(c[0]), stringToNumbers(c[1])}
}

func stringToNumbers(s string) []int {
	gs := strings.Fields(s)

	var res []int

	for _, v := range gs {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Errorf("error: %v", err)
		}

		res = append(res, n)
	}

	return res
}

func Solve2() any {
	lns := utils.ReadLines("04.txt")

	m := make(map[int]int)
	total := 0

	for i, ln := range lns {
		card := loadS(ln)

		w := make(map[int]bool)
		for _, v := range card.W {
			w[v] = true
		}

		res := 0

		for _, v := range card.S {
			if w[v] {
				res++
			}
		}

		for j := i + 1; j <= i+res; j++ {
			m[j] = m[j] + m[i] + 1
		}
		total = i
	}

	ans := 0

	for i := 0; i <= total; i++ {
		ans += m[i] + 1
	}
	return ans
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)
}
