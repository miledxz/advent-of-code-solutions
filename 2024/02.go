package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/miledxz/advent-of-code-solutions/utils"
)

func Solve1() int {
	file, err := os.ReadFile("02.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileString := string(file)
	grid := convInput(fileString)

	var ans int
	for _, lvl := range grid {
		if compare(lvl) {
			ans++
		}
	}
	return ans
}

func Solve2() int {
	file, err := os.ReadFile("02.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileString := string(file)
	grid := convInput(fileString)

	var ans int
	for _, lvl := range grid {
		for i := range len(lvl) {
			newLvl := []int{}
			for j := range len(lvl) {
				if i != j {
					newLvl = append(newLvl, lvl[j])
				}
			}
			if compare(newLvl) {
				ans++
				break
			}
		}
	}

	return ans

}

func compare(level []int) bool {
	isIncreasing := level[1] > level[0]

	for i := 1; i < len(level); i++ {
		if isIncreasing && level[i] <= level[i-1] {
			return false
		} else if !isIncreasing && level[i] >= level[i-1] {
			return false
		}

		diff := utils.Abs(level[i] - level[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func convInput(input string) [][]int {
	grid := [][]int{}
	for _, ln := range strings.Split(input, "\n") {
		level := []int{}
		for _, n := range strings.Split(ln, " ") {
			level = append(level, utils.ToInt(n))
		}
		grid = append(grid, level)
	}
	return grid
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)
}
