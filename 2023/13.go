package main

import (
	"fmt"

	"github.com/miledxz/advent-of-code-solutions/utils"
)

func Solve1() any {
	lns := utils.ReadLines("13.txt")
	start := 0

	sum := 0
	for i, v := range lns {
		if v == "" || i == len(lns)-1 {
			f := parseField(lns[start:i])

			v := sym1(f.V)
			sum += v
			h := sym1(f.H)
			sum += 100 * h

			start = i + 1
		}
	}

	return sum
}

func Solve2() any {
	lns := utils.ReadLines("13.txt")
	start := 0

	sum := 0
	for i, v := range lns {
		if v == "" || i == len(lns)-1 {
			f := parseField(lns[start:i])

			v := sym2(f.V)
			sum += v
			h := sym2(f.H)
			sum += 100 * h

			if v != 0 && h != 0 {
				fmt.Println(v, h)
				for l := start; l < i; l++ {
					fmt.Println(lns[l])
				}
				fmt.Println(f.V)
				fmt.Println(f.H)
			}
			start = i + 1
		}
	}

	return sum
}

type Field struct {
	H []uint64
	V []uint64
}

func parseField(lns []string) Field {
	hz := []uint64{}
	vz := []uint64{}

	for j, ln := range lns {
		var hval uint64
		for i, r := range ln {
			var c uint64
			if r == '#' {
				c = 1
			}

			hval |= c << i
			if j == 0 {
				vz = append(vz, c)
			} else {
				vz[i] |= c << j
			}
		}
		hz = append(hz, hval)
	}

	return Field{hz, vz}
}

func sym1(lines []uint64) int {
	for i := 0; i < len(lines)-1; i++ {
		for c := 0; c < len(lines)/2+1; c++ {
			l := i - c
			r := i + c + 1

			if l < 0 || r >= len(lines) {
				return i + 1
			}

			if lines[l] != lines[r] {
				break
			}
		}
	}

	return 0
}

func sym2(lns []uint64) int {
	for i := 0; i < len(lns)-1; i++ {
		offByOne := false
		for c := 0; c < len(lns)/2+1; c++ {
			l := i - c
			r := i + c + 1

			if l < 0 || r >= len(lns) {
				if offByOne {
					return i + 1
				} else {
					continue
				}
			}

			if lns[l] != lns[r] {
				difference := lns[l] ^ lns[r]
				if utils.CountBits(difference) == 1 && !offByOne {
					offByOne = true
				} else {
					break
				}
			}
		}
	}

	return 0
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)

}
