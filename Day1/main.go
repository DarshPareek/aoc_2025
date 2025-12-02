package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
func turnLeft(dial int, rot int) (int, int) {
	zs := 0
	for i := 0; i < rot; i++ {
		dial -= 1
		// Handle wrapping first
		if dial == -1 {
			dial = 99
		}
		// Then check if we landed on 0
		if dial == 0 {
			zs += 1
		}
	}
	return dial, zs
}

func turnRight(dial int, rot int) (int, int) {
	zs := 0
	for i := 0; i < rot; i++ {
		dial += 1
		// Handle wrapping first
		if dial == 100 {
			dial = 0
		}
		// Then check if we landed on 0
		if dial == 0 {
			zs += 1
		}
	}
	return dial, zs
}

func main() {
	filepath := "./input.txt"
	dial := 50
	zs := 0
	content, err := readInput(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	for _, value := range content {
		fmt.Println(dial)
		dir := string(value[0])
		n, err := strconv.Atoi(string(value[1:]))
		if err != nil {
			fmt.Printf("Error converting string to int: %v", err)
		}
		if dir == "L" {
			d, z := turnLeft(dial, n)
			dial = d
			zs += z
		} else {
			d, z := turnRight(dial, n)
			dial = d
			zs += z
		}
	}

	fmt.Println("Hello, World!", zs)
}
