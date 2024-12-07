package main

import (
	"/utils"
	"fmt"
)

func Solve1() int64 {
	lns := utils.ReadLines("03.txt")

	el := make(map[utils.Point]rune)

	nums := []Num{}

	for j, ln := range lns {
		numint := false

		for i, r := range ln {
			switch {
			case r == '.':
				numint = false
			case r >= '0' && r <= '9':
				switch numint {
				case false:
					numint = true
					var num Num
					num.S = utils.Point{X: i, Y: j}
					num.E = num.S
					num.V = int64(r - '0')
					nums = append(nums, num)
				case true:
					idx := len(nums) - 1
					nums[idx].E = utils.Point{X: i, Y: j}
					nums[idx].V = nums[idx].V*10 + int64(r-'0')
				}
			default:
				el[utils.Point{X: i, Y: j}] = r
				numint = false
			}
		}
	}

	var sum int64

	for _, n := range nums {
		adjP := ajdPoints(n.S, n.E)
		for _, p := range adjP {
			if el[p] != 0 {
				sum += n.V
				break
			}
		}
	}

	return sum
}

type Num struct {
	V int64
	S utils.Point
	E utils.Point
}

func ajdPoints(start utils.Point, end utils.Point) []utils.Point {
	yu := start.Y - 1
	yd := start.Y + 1
	yl := start.Y

	xmin := start.X - 1
	xmax := end.X + 1

	res := []utils.Point{{X: xmin, Y: yl}, {X: xmax, Y: yl}}

	for i := xmin; i <= xmax; i++ {
		res = append(res, utils.Point{X: i, Y: yu})
		res = append(res, utils.Point{X: i, Y: yd})
	}

	return res
}

func Solve2() any {
	lns := utils.ReadLines("03.txt")

	el := make(map[utils.Point]rune)

	nums := []Num{}

	for j, ln := range lns {
		numint := false

		for i, r := range ln {
			switch {
			case r == '.':
				numint = false
			case r >= '0' && r <= '9':
				switch numint {
				case false:
					numint = true
					var num Num
					num.S = utils.Point{X: i, Y: j}
					num.E = num.S
					num.V = int64(r - '0')
					nums = append(nums, num)
				case true:
					idx := len(nums) - 1
					nums[idx].E = utils.Point{X: i, Y: j}
					nums[idx].V = nums[idx].V*10 + int64(r-'0')
				}
			default:
				el[utils.Point{X: i, Y: j}] = r
				numint = false
			}
		}
	}

	var sum int64

	adjG := make(map[utils.Point][]int64)

	for _, n := range nums {
		adjP := ajdPoints(n.S, n.E)
		for _, p := range adjP {
			if el[p] == '*' {
				adjG[p] = append(adjG[p], n.V)
				break
			}
		}
	}

	for _, numl := range adjG {
		if len(numl) == 2 {
			sum += numl[0] * numl[1]
		}
	}

	return sum
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)
}
