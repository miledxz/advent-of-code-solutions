package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/miledxz/advent-of-code-solutions/utils"
)

func Solve1() int {
	lns := utils.ReadLines("01.txt")

	var l, r []int

	for _, ln := range lns {
		nums := strings.Split(ln, "   ")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		l = append(l, num1)
		r = append(r, num2)

	}

	sort.Ints(l)
	sort.Ints(r)

	var res int
	for i := 0; i < len(l); i++ {
		res += abs(l[i] - r[i])
	}

	return res
}

func Solve2() int {
	lns := utils.ReadLines("01.txt")

	var l, r []int

	for _, ln := range lns {
		nums := strings.Split(ln, "   ")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		l = append(l, num1)
		r = append(r, num2)

	}

	m := make(map[int]int)

	for _, v1 := range l {
		for _, v2 := range r {
			if v1 == v2 {
				m[v1]++
			}
		}
	}

	var res int
	for _, v := range l {
		res += v * m[v]
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)
}
