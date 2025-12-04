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
func convertToMap(lines []string) [][]string {
	var res [][]string
	for i := 0; i < len(lines); i++ {
		var temp []string
		for j := 0; j < len(lines[i]); j++ {
			temp = append(temp, string(lines[i][j]))
		}
		res = append(res, temp)
	}
	return res
}
func main() {
	lines, err := readInput("./input.txt")
	if err != nil {
		fmt.Println("Good Bye")
		return
	}
	var neighbors [8][2]int
	neighbors = [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	n := len(lines)
	mov := 1
	new_mov := 0
	grid := convertToMap(lines)
	for mov != 0 {
		mov = 0
		for i := 0; i < n; i++ {
			l := lines[i]
			m := len(l)
			for j := 0; j < m; j++ {
				if string(grid[i][j]) == "@" {
					nei := 0
					for k := 0; k < 8; k++ {
						nx := j + neighbors[k][0]
						ny := i + neighbors[k][1]
						if nx > -1 && nx < m && ny > -1 && ny < m {
							if string(grid[ny][nx]) == "@" {
								nei += 1
							}
						}
					}

					if nei < 4 {
						fmt.Println(i, j)
						grid[i][j] = "."
						mov += 1
					}
				}
			}
		}
		new_mov += mov
	}
	fmt.Println(mov, new_mov)
}
