package main

import (
	"fmt"
	"strings"

	"aoc/utils"
)

func Solve1() int {
	lns := utils.ReadLines("01.txt")

	var l, r int
	ans := 0

	for _, ln := range lns {
		l = -1
		r = -1
		for _, el := range ln {
			if el >= '0' && el <= '9' {
				if l == -1 {
					l = int(el - '0')
				}
				r = int(el - '0')
			}
		}
		ans += 10*l + r
	}
	return ans
}

func Solve2() int {
	lns := utils.ReadLines("01.txt")

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	ans := 0

	for _, ln := range lns {
		l := -1
		r := -1
		min := len(ln) + 1
		max := -1

		for k, v := range m {
			p := strings.Index(ln, k)
			if p == -1 {
				continue
			}
			if p < min {
				l = v
				min = p
			}

			p = strings.LastIndex(ln, k)
			if p > max {
				r = v
				max = p
			}
		}
		ans += 10*l + r
	}
	return ans
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)
}
