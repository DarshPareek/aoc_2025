package main

import (
	"bufio"
	"fmt"
	"os"
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

func ini(lines []string) []int {
	var state []int
	for i := 0; i < len(lines[0]); i++ {
		if string(lines[0][i]) != "S" {
			state = append(state, 0)
		} else {
			state = append(state, 1)
		}
	}
	return state
}
func sum(state []int) int {
	s := 0
	for i := 0; i < len(state); i++ {
		s += state[i]
	}
	return s
}

func main() {
	lines, _ := readInput("./input.txt")
	res := 0
	state := ini(lines)
	var timelines [][]int
	timelines = append(timelines, state)
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			if string(line[j]) == "^" && state[j] > 0 {
				res += 1
				state[j-1] += timelines[i-1][j]
				state[j+1] += timelines[i-1][j]
				state[j] = 0
			}
		}
		timelines = append(timelines, state)
		fmt.Println(sum(state), i, state)
	}
	fmt.Println(res, state, sum(state))
}
