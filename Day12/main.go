package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("error reading the file %v\n", err)
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func sum(arr []int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}
func main() {
	lines, _ := readInput("./input.txt")
	shapes := lines[:30]
	regions := lines[30:]
	var shape_size []int
	for i := 0; i < len(shapes); i++ {
		if i%5 == 0 || i == 0 {
			continue
		}
		t := 0
		for ; shapes[i] != ""; i++ {
			for j := 0; j < len(shapes[i]); j++ {
				if string(shapes[i][j]) == "#" {
					t += 1
				}
			}
		}
		shape_size = append(shape_size, t)
	}
	fmt.Println(shape_size)
	res := 0
	for i := 0; i < len(regions); i++ {
		var final_p []int
		t := strings.Split(regions[i], ":")
		size_data := t[0]
		p_data := strings.Fields(t[1])
		s := strings.Split(size_data, "x")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		gs := x * y
		for i := 0; i < len(p_data); i++ {
			t, _ := strconv.Atoi(p_data[i])
			final_p = append(final_p, t*shape_size[i])
		}
		ps := sum(final_p)
		if int(float64(ps)*1.3) < gs {
			res += 1
		} else if ps > gs {
			continue
		} else {
			fmt.Println("Not okay")
		}
	}
	fmt.Println(res)
}
