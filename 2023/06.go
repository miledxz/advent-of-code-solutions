package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/miledxz/advent-of-code-solutions/utils"
)

func Solve1() any {
	lns := utils.ReadLines("06.txt")

	ts := strings.Fields(lns[0])
	ds := strings.Fields(lns[1])

	sum := 1

	for i := 1; i < len(ts); i++ {
		time, err := strconv.ParseFloat(ts[i], 64)
		if err != nil {
			fmt.Errorf("error: %v", err)
		}

		dist, err := strconv.ParseFloat(ds[i], 64)
		if err != nil {
			fmt.Errorf("error: %v", err)
		}

		tpMin := 0.5*time - 0.5*math.Sqrt(time*time-4.0*dist)

		tpMax := 0.5*time + 0.5*math.Sqrt(time*time-4.0*dist)

		ww := math.Floor(tpMax) - math.Ceil(tpMin)
		if math.Abs(tpMin-math.Ceil(tpMin)) < 1.0e-8 {
			ww -= 1.0
		}

		if math.Abs(tpMax-math.Floor(tpMax)) < 1.0e-8 {
			ww -= 1.0
		}

		sum *= int(ww) + 1
	}
	return sum
}

func Solve2() any {
	lns := utils.ReadLines("06.txt")

	ts := strings.Fields(lns[0])
	ds := strings.Fields(lns[1])

	timeString := strings.Join(ts[1:], "")
	distanceString := strings.Join(ds[1:], "")

	time, err := strconv.ParseFloat(timeString, 64)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	dist, err := strconv.ParseFloat(distanceString, 64)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	tpMin := 0.5*time - 0.5*math.Sqrt(time*time-4.0*dist)

	tpMax := 0.5*time + 0.5*math.Sqrt(time*time-4.0*dist)

	ww := math.Floor(tpMax) - math.Ceil(tpMin)
	if math.Abs(tpMin-math.Ceil(tpMin)) < 1.0e-8 {
		ww -= 1.0
	}

	if math.Abs(tpMax-math.Floor(tpMax)) < 1.0e-8 {
		ww -= 1.0
	}

	return int(ww) + 1
}

func main() {
	ans1 := Solve1()
	ans2 := Solve2()
	fmt.Println(ans1)
	fmt.Println(ans2)

}
