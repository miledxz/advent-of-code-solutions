package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Solve1() any {
	lns := utils.ReadLines("08.txt")

	moves := []rune(lns[0])

	adj := make(map[string][]string)
	for i := 2; i < len(lns); i++ {
		fs := strings.Split(lns[i], " = ")

		loc := strings.Trim(fs[1], "()")

		lr := strings.Split(loc, ", ")

		adj[fs[0]] = lr
	}

	next := adj["AAA"]
	steps := 0

outer:
	for {
		for _, d := range moves {
			steps++
			switch d {
			case 'L':
				if next[0] == "ZZZ" {
					break outer
				}
				next = adj[next[0]]
			case 'R':
				if next[1] == "ZZZ" {
					break outer
				}
				next = adj[next[1]]
			}
		}
	}

	return steps
}

func Solve2() any {
	lns := utils.ReadLines("08.txt")

	moves := []rune(lns[0])

	adj := make(map[string][]string)
	for i := 2; i < len(lns); i++ {
		fs := strings.Split(lns[i], " = ")

		loc := strings.Trim(fs[1], "()")

		lr := strings.Split(loc, ", ")

		adj[fs[0]] = lr
	}

	cyclesteps := int64(1)

	for pos := range adj {
		if !strings.HasSuffix(pos, "A") {
			continue
		}
		var steps int64

		hits := []Hit{}

	outer:
		for {
			for _, d := range moves {
				next := adj[pos]
				steps++
				switch d {
				case 'L':
					pos = next[0]
				case 'R':
					pos = next[1]
				}

				if strings.HasSuffix(pos, "Z") {
					h := Hit{pos, steps}
					for _, ph := range hits {
						if h.El == ph.El && h.S == 2*ph.S {
							cyclesteps = utils.Lcm(cyclesteps, h.S-ph.S)
							break outer
						}
					}
					hits = append(hits, h)
				}
			}
		}
	}
	return cyclesteps
}

type Hit struct {
	El string
	S     int64
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)
}
