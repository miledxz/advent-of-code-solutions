package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed 03.txt
var input string

func main() {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	ans1 := 0
	ans2 := 0
	enabled := true
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	for _, line := range lines {
		for _, match := range r.FindAllString(line, -1) {
			if match == "do()" {
				enabled = true
			} else if match == "don't()" {
				enabled = false
			} else {
				s := strings.Split(match[4:len(match)-1], ",")
				x, _ := strconv.Atoi(s[0])
				y, _ := strconv.Atoi(s[1])
				ans1 += x * y
				if enabled {
					ans2 += x * y
				}
			}
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}
