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

func solveWorksheet(worksheet [][]int) int {
	var res int
	start := worksheet[0]
	op := worksheet[len(worksheet)-1]
	for i := 1; i < len(worksheet)-1; i++ {
		for j := 0; j < len(start); j++ {
			fmt.Println("Going to ", op[j], worksheet[i][j], "and", start[j])
			if op[j] == 0 {
				worksheet[i][j] *= start[j]
			}
			if op[j] == 1 {
				worksheet[i][j] += start[j]
			}
		}
		start = worksheet[i]
	}
	for i := 0; i < len(worksheet[len(worksheet)-2]); i++ {
		res += worksheet[len(worksheet)-2][i]
	}
	return res
}

func compress(worksheet [][]int, col int) int {
	var res int
	p := ""
	for i := 0; i < len(worksheet)-1; i++ {
		if worksheet[i][col] != 0 {
			p += strconv.Itoa(worksheet[i][col])
		}
	}
	res, _ = strconv.Atoi(p)
	fmt.Println(res, col)
	return res
}

func main() {
	lines, _ := readInput("./input.txt")
	var worksheet [][]int
	for i := 0; i < len(lines)-1; i++ {
		temp := lines[i] //strings.Split(lines[i], "")
		fmt.Println(temp)
		var tn []int
		for i := 0; i < len(temp); i++ {
			n, _ := strconv.Atoi(string(temp[i]))
			tn = append(tn, n)
		}
		worksheet = append(worksheet, tn)
	}
	fmt.Println(worksheet)
	var tn []int
	temp := strings.Split(lines[len(lines)-1], "")
	fmt.Println(temp)
	p := 0
	if string(temp[0]) == "+" {
		p = 1
	}
	for i := 0; i < len(temp); i++ {
		if string(temp[i]) == "+" {
			p = 1
		} else if string(temp[i]) == "*" {
			p = 0
		}
		tn = append(tn, p)
	}
	worksheet = append(worksheet, tn)
	op := worksheet[len(worksheet)-1]
	for i := 0; i < len(worksheet); i++ {
		worksheet[i] = append(worksheet[i], -1)
	}
	fmt.Println(worksheet)
	for i := 0; i < len(worksheet); i++ {
		fmt.Println(worksheet[i])
	}
	var m [][]int
	var nn []int
	for i := 0; i <= len(op); i++ {
		res := compress(worksheet, i)
		nn = append(nn, res)
		if res == 0 {
			nn = append(nn, op[i-1])
			m = append(m, nn)
			nn = nil
			fmt.Println("All prev will be", op[i-1])
		}
	}
	fmt.Println(m)
	var res int
	res = 0
	for i := 0; i < len(m); i++ {
		var add int
		var mul int
		add = 0
		mul = 1
		for j := 0; j < len(m[i])-1; j++ {
			op := m[i][len(m[i])-1]
			if op == 0 && m[i][j] != 0 {
				mul *= m[i][j]
			} else if op == 1 && m[i][j] != 0 {
				add += m[i][j]
			}
		}
		if mul == 1 && add != 0 {
			res += add
		} else {
			res += mul
		}
	}
	fmt.Println(res)
	//res := solveWorksheet(worksheet)
	//fmt.Println(res)
}
