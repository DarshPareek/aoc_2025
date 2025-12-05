package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Range struct makes handling start/end pairs easier to read
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

	// Split ranges and ingredients based on the empty line
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

	// Fallback if no empty line found (all ranges)
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

// part2 calculates the number of unique integers covered by the ranges
func part2(rawRanges []string) int {
	ranges := parseRanges(rawRanges)

	// 1. Sort ranges by start value
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	if len(ranges) == 0 {
		return 0
	}

	// 2. Merge overlapping intervals
	var merged []Range
	merged = append(merged, ranges[0])

	for i := 1; i < len(ranges); i++ {
		current := ranges[i]
		lastMerged := &merged[len(merged)-1]

		// Check for overlap.
		// Since we sorted by Start, we only need to check if current.Start <= lastMerged.End
		// We use `+1` here to merge touching ranges (e.g. 3-4 and 5-6 become 3-6),
		// though strictly for counting it doesn't change the math, it keeps the list cleaner.
		if current.Start <= lastMerged.End+1 {
			// If the current range goes further than the previous one, extend the previous one
			if current.End > lastMerged.End {
				lastMerged.End = current.End
			}
		} else {
			// No overlap, start a new range
			merged = append(merged, current)
		}
	}

	// 3. Count the total integers in the merged, disjoint ranges
	count := 0
	for _, r := range merged {
		count += (r.End - r.Start) + 1
	}

	return count
}

// part1 (kept largely the same but cleaned up slightly)
func part1(rawRanges []string, ing []string) int {
	// Parse ranges once for performance
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
	// Ensure input.txt exists in the same directory
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
