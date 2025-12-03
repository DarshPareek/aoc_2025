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

func stringToArr(arr string) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		d, err := strconv.Atoi(string(arr[i]))
		res = append(res, d)
		if err != nil {
			fmt.Println("Not Possible")
			return nil
		}
	}
	return res
}

func findMax(arr []int) (int, int) {
	res := arr[0]
	resi := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
			resi = i
		}
	}
	return res, resi
}
func findSecondMax(arr []int, ind int) (int, int) {
	res := 0
	resi := -1
	for i := ind + 1; i < len(arr); i++ {
		if arr[i] > res && i != ind {
			res = arr[i]
			resi = i
		}
	}
	return res, resi
}

func formJolt(r, s, i, j int) int {
	if i > j {
		return s*10 + r
	}
	return r*10 + s
}

func arrToNum(arr []int, arr2 []int, ind int) int {
	res := ""
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 || i == ind {
			res += strconv.Itoa(arr2[i])
		}
	}
	c, _ := strconv.Atoi(res)
	return c
}

func makeBestChoice(arr []int, best int, ind int, used []int) (int, int, []int) {
	resi := 0
	res := best
	for i := 0; i < len(arr); i++ {
		cs := strconv.Itoa(best)
		if cs == "0" {
			cs = ""
		}
		if used[i] == 0 {
			//var num1 int
			//var num2 int
			//var num int
			//if ind < i {
			//	cs += strconv.Itoa(arr[i])
			//	c, _ := strconv.Atoi(cs)
			//	num1 = c
			//} else {
			//	cs = strconv.Itoa(arr[i]) + cs
			//	c, _ := strconv.Atoi(cs)
			//	num2 = c
			//}
			//if num1 > num2 {
			//	num = num1
			//} else {
			//	num = num2
			//}
			num := arrToNum(used, arr, i)
			//fmt.Println(num)
			if num > res {
				res = num
				resi = i
			}
		}
	}
	used[resi] = arr[resi]
	return res, resi, used
}

func ini(arr []int) []int {
	var used []int
	for i := 0; i < len(arr); i++ {
		used = append(used, 0)
	}
	return used
}

func main() {
	data, err := readInput("./input.txt")
	if err != nil {
		fmt.Println("Failed")
		return
	}
	//var dp map[int][]int
	var res int
	for i := 0; i < len(data); i++ {
		t := stringToArr(data[i])
		used := ini(t)
		r, i := findMax(t)
		used[i] = r
		curr, ind, used := makeBestChoice(t, r, i, used)
		curr, ind, used = makeBestChoice(t, curr, ind, used)
		curr, ind, used = makeBestChoice(t, curr, ind, used)
		curr, ind, used = makeBestChoice(t, curr, ind, used)
		curr, ind, used = makeBestChoice(t, curr, ind, used)
		curr, ind, used = makeBestChoice(t, curr, ind, used)
		curr, ind, used = makeBestChoice(t, curr, ind, used)
		curr, ind, used = makeBestChoice(t, curr, ind, used)
		curr, ind, used = makeBestChoice(t, curr, ind, used)
		curr, ind, used = makeBestChoice(t, curr, ind, used)
		curr, ind, used = makeBestChoice(t, curr, ind, used)
		//curr, ind, used = makeBestChoice(t, curr, ind, used)
		fmt.Println(curr, ind, used)
		res += curr
		//r, i := findMax(t)
		//var fh []int
		//var s int
		//var j int
		//var num int
		//if i == 0 {
		//	fh = t[i+1:]
		//	s, j = findMax(fh)
		//	num = formJolt(r, s, i, j)
		//} else if i == len(t)-1 {
		//	fh = t[:i]
		//	s, j = findMax(fh)
		//	num = formJolt(r, s, i, j)
		//} else {
		//	fh = t[:i]
		//	s, j = findMax(fh)
		//	sh := t[i+1:]
		//	t, k := findMax(sh)
		//	num1 := formJolt(r, s, i, j)
		//	num2 := formJolt(r, t, i, k+i)
		//	if num1 > num2 {
		//		num = num1
		//	} else {
		//		num = num2
		//	}
		//}
		//res += num
	}
	fmt.Println(res)
}
