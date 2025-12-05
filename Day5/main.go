package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func readInput(filepath string) ([]string, []string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	var ranges []string
	var ing []string
	for i, line := range lines {
		if line == "" {
			ranges = lines[:i]
			if i+1 < len(lines) {
				ing = lines[i+1:]
			}
			break
		}
	}
	if len(ranges) == 0 && len(ing) == 0 {
		ranges = lines
	}

	return ranges, ing, nil
}

func parseRanges(rawRanges []string) []Range {
	var ranges []Range
	for _, r := range rawRanges {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}
		lb, _ := strconv.Atoi(parts[0])
		ub, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{Start: lb, End: ub})
	}
	return ranges
}

func part2(rawRanges []string) int {
	ranges := parseRanges(rawRanges)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
	if len(ranges) == 0 {
		return 0
	}
	p2 := 0
	current := -1
	for i := 0; i < len(ranges); i++ {
		if current >= ranges[i].Start {
			ranges[i].Start = current + 1
		}
		if ranges[i].Start <= ranges[i].End {
			p2 += ranges[i].End - ranges[i].Start + 1
		}
		current = max(ranges[i].End, current)
	}

	return p2
}

func part1(rawRanges []string, ing []string) int {
	ranges := parseRanges(rawRanges)
	res := 0
	for _, rawN := range ing {
		n, _ := strconv.Atoi(rawN)
		isFresh := false
		for _, r := range ranges {
			if n >= r.Start && n <= r.End {
				isFresh = true
				break
			}
		}
		if isFresh {
			res++
		}
	}
	return res
}

func main() {
	ranges, ing, err := readInput("./input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	p1 := part1(ranges, ing)
	p2 := part2(ranges)

	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2 (Total Fresh IDs): %d\n", p2)
}
