package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Solve1() int64 {
	lns := utils.ReadLines("05.txt")
	p := strings.Fields(lns[0])
	vs := make([]int64, len(p)-1)

	for i := 0; i < len(vs); i++ {
		v, err := strconv.ParseInt(p[i+1], 10, 64)
		if err != nil {
			fmt.Errorf("error: %v", err)
		}

		vs[i] = v
	}

	diff := []Diff{}

	re := regexp.MustCompile(`\d+\s+\d+\s+\d+`)

	for _, ln := range lns[2:] {
		switch {
		case re.MatchString(ln):
			p = strings.Fields(ln)

			els := make([]int64, len(p))

			for i, s := range p {
				var err error
				els[i], err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					fmt.Errorf("error: %v", err)
				}
			}
			c := Diff{S: els[1], E: els[1] + els[2], Move: els[0] - els[1]}
			diff = append(diff, c)

		case ln == "":
			sort.Slice(diff, func(i, j int) bool { return diff[i].S < diff[j].S })

			for i, v := range vs {
				move := getDelta(diff, v)
				vs[i] = v + move
			}

			diff = []Diff{}
		}
	}

	if len(diff) > 0 {
		sort.Slice(diff, func(i, j int) bool { return diff[i].S < diff[j].S })

		for i, v := range vs {
			move := getDelta(diff, v)
			vs[i] = v + move
		}
	}

	min := vs[0]

	for _, v := range vs[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

type Diff struct {
	S int64
	E   int64
	Move int64
}

func getDelta(arr []Diff, val int64) int64 {
	l, r := 0, len(arr)-1

	for l <= r {
		m := l + (r-l)/2

		if arr[m].S <= val && arr[m].E > val {
			return arr[m].Move
		}

		if arr[m].S > val {
			r = m - 1
		} else if arr[m].E <= val {
			l = m + 1
		}
	}

	return 0
}

func Solve2() int64 {
	lns := utils.ReadLines("05.txt")

	p := strings.Fields(lns[0])

	vs := make([]int64, len(p)-1)

	for i := 0; i < len(vs); i++ {
		v, err := strconv.ParseInt(p[i+1], 10, 64)
		if err != nil {
			fmt.Errorf("error: %v", err)
		}

		vs[i] = v
	}

	diff := []Diff{}

	re := regexp.MustCompile(`\d+\s+\d+\s+\d+`)

	for _, ln := range lns[2:] {
		switch {
		case re.MatchString(ln):
			p = strings.Fields(ln)

			els := make([]int64, len(p))

			for i, s := range p {
				var err error
				els[i], err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					fmt.Errorf("error: %v", err)
				}
			}
			c := Diff{S: els[1], E: els[1] + els[2], Move: els[0] - els[1]}
			diff = append(diff, c)

		case ln == "":
			vs = diffCalc(diff, vs)

			diff = []Diff{}
		}
	}

	if len(diff) > 0 {
		vs = diffCalc(diff, vs)
	}

	min := vs[0]

	for i := 2; i < len(vs); i += 2 {
		if vs[i] < min {
			min = vs[i]
		}
	}

	return min
}

func diffMove(arr []Diff, val int64) (int64, int64) {
	l, r := 0, len(arr)-1
	var m int

	for l <= r {
		m = l + (r-l)/2

		if arr[m].S <= val && arr[m].E > val {
			return arr[m].Move, arr[m].E - val
		}

		if arr[m].S > val {
			r = m - 1
		} else if arr[m].E <= val {
			l = m + 1
		}
	}

	if arr[m].S > val {
		return 0, arr[m].S - val
	} else if m < len(arr)-1 {
		return 0, arr[m+1].S - val
	} else {
		return 0, 0
	}
}

func diffCalc(diff []Diff, vs []int64) []int64 {
	sort.Slice(diff, func(i, j int) bool { return diff[i].S < diff[j].S })
	vls := []int64{}

	for i := 0; i < len(vs); i += 2 {
		start, length := vs[i], vs[i+1]

		for {
			move, diff := diffMove(diff, start)
			vls = append(vls, start+move)
			if length <= diff || diff == 0 {
				vls = append(vls, length)
				break
			} else {
				vls = append(vls, diff)
				start = start + diff
				length = length - diff
			}
		}
	}
	return vls
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)

}
