package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type File struct {
	Pos    int
	Size   int
	FileID int
}

type Dash struct {
	Pos  int
	Size int
}

func Solve(input string, part2 bool) int {
	var file []File
	var dash []Dash
	var res []interface{}
	fileID := 0
	pos := 0

	for i, c := range input {
		count, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			if part2 {
				file = append(file, File{Pos: pos, Size: count, FileID: fileID})
			}
			for j := 0; j < count; j++ {
				res = append(res, fileID)
				if !part2 {
					file = append(file, File{Pos: pos, Size: 1, FileID: fileID})
				}
				pos++
			}
			fileID++
		} else {
			dash = append(dash, Dash{Pos: pos, Size: count})
			for j := 0; j < count; j++ {
				res = append(res, nil)
				pos++
			}
		}
	}

	for i := len(file) - 1; i >= 0; i-- {
		file := file[i]
		for j, space := range dash {
			if space.Pos < file.Pos && file.Size <= space.Size {
				for k := 0; k < file.Size; k++ {
					if res[file.Pos+k] != file.FileID {
						panic(fmt.Sprintf("assertion err: res[%d] != fileID", file.Pos+k))
					}
					res[file.Pos+k] = nil
					res[space.Pos+k] = file.FileID
				}
				dash[j] = Dash{Pos: space.Pos + file.Size, Size: space.Size - file.Size}
				break
			}
		}
	}

	checksum := 0
	for i, c := range res {
		if c != nil {
			checksum += i * c.(int)
		}
	}

	return checksum
}

func Solve1(input string) int {
	return Solve(input, false)
}

func Solve2(input string) int {
	return Solve(input, true)
}

func main() {
	fn := "09.txt"
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	filePath := filepath.Join(cwd, fn)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	result1 := Solve1(input)
	result2 := Solve2(input)

	fmt.Println(result1)
	fmt.Println(result2)
}
