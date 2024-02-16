package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

var (
	maxr = 12
	maxg = 13
	maxb = 14
)

func Solve1() int {
	lns := utils.ReadLines("02.txt")
	ans := 0
o:
	for _, ln := range lns {
		g := loadG(ln)
		for _, d := range g.h {
			if d["red"] > maxr || d["green"] > maxg || d["blue"] > maxb {
				continue o
			}
		}
		ans += g.id
	}

	return ans
}

type Game struct {
	id int
	h  []map[string]int
}

func loadG(ln string) Game {
	ans := Game{}
	fs := strings.Split(ln, ": ")
	ids := strings.Fields(fs[0])
	var err error
	ans.id, err = strconv.Atoi(ids[1])
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	ds := strings.Split(fs[1], "; ")

	for _, d := range ds {
		dm := make(map[string]int)
		cs := strings.Split(d, ", ")
		for _, s := range cs {
			p := strings.Fields(s)
			n, err := strconv.Atoi(p[0])
			if err != nil {
				fmt.Errorf("error: %v", err)
			}
			dm[p[1]] = n
		}
		ans.h = append(ans.h, dm)
	}

	return ans
}

func Solve2() int {
	lns := utils.ReadLines("02.txt")

	ans := 0

	for _, ln := range lns {
		g := loadG(ln)

		maxm := make(map[string]int)
		for _, d := range g.h {
			for k, v := range d {
				if v > maxm[k] {
					maxm[k] = v
				}
			}
		}

		ans += maxm["red"] * maxm["green"] * maxm["blue"]
	}

	return ans
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)

}
