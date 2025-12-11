package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parseLine extracts the buttons (list of []int slices) and targets ([]int) from one input line.
// It ignores the indicator diagram in [brackets].
func parseLine(line string) ([][]int, []int, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil, nil
	}

	var buttons [][]int
	var targets []int

	// Walk the line and pull out (...) and {...} segments.
	for i := 0; i < len(line); {
		switch line[i] {
		case '[':
			// skip until matching ]
			j := strings.IndexByte(line[i:], ']')
			if j < 0 {
				return nil, nil, fmt.Errorf("unmatched [")
			}
			i += j + 1
		case '(':
			j := strings.IndexByte(line[i:], ')')
			if j < 0 {
				return nil, nil, fmt.Errorf("unmatched (")
			}
			content := line[i+1 : i+j]
			content = strings.TrimSpace(content)
			if content == "" {
				// empty button: no indices -> ignore
			} else {
				parts := strings.Split(content, ",")
				var idxs []int
				for _, p := range parts {
					p = strings.TrimSpace(p)
					if p == "" {
						continue
					}
					n, err := strconv.Atoi(p)
					if err != nil {
						return nil, nil, fmt.Errorf("bad button index %q", p)
					}
					idxs = append(idxs, n)
				}
				buttons = append(buttons, idxs)
			}
			i += j + 1
		case '{':
			j := strings.IndexByte(line[i:], '}')
			if j < 0 {
				return nil, nil, fmt.Errorf("unmatched {")
			}
			content := line[i+1 : i+j]
			content = strings.TrimSpace(content)
			if content != "" {
				parts := strings.Split(content, ",")
				for _, p := range parts {
					p = strings.TrimSpace(p)
					if p == "" {
						continue
					}
					n, err := strconv.Atoi(p)
					if err != nil {
						return nil, nil, fmt.Errorf("bad target %q", p)
					}
					targets = append(targets, n)
				}
			}
			i += j + 1
		default:
			i++
		}
	}

	return buttons, targets, nil
}

// encodeState makes a simple string key for a counter-state.
func encodeState(state []int) string {
	parts := make([]string, len(state))
	for i, v := range state {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, ",")
}

// minPressesForMachine uses BFS to find the fewest button presses to reach targets exactly.
// Returns -1 if unreachable.
func minPressesForMachine(buttons [][]int, targets []int) int {
	n := len(targets)
	// start state is all zeros
	start := make([]int, n)
	targetKey := encodeState(targets)

	// Quick check: if no counters (n == 0) then zero presses needed.
	if n == 0 {
		return 0
	}

	// Pre-filter buttons: remove any that affect indices outside range or that affect nothing.
	validButtons := [][]int{}
	for _, b := range buttons {
		valid := true
		if len(b) == 0 {
			// button affects nothing: skip (useless)
			continue
		}
		for _, idx := range b {
			if idx < 0 || idx >= n {
				valid = false
				break
			}
		}
		if valid {
			validButtons = append(validButtons, b)
		}
	}
	if len(validButtons) == 0 {
		// no valid button can change counters; reachable only if all targets are zero
		for _, t := range targets {
			if t != 0 {
				return -1
			}
		}
		return 0
	}

	type node struct {
		state []int
		dist  int
	}
	q := make([]node, 0, 1024)
	visited := make(map[string]struct{})

	q = append(q, node{state: start, dist: 0})
	visited[encodeState(start)] = struct{}{}

	head := 0
	for head < len(q) {
		cur := q[head]
		head++

		curKey := encodeState(cur.state)
		if curKey == targetKey {
			return cur.dist
		}

		// Try pressing each button once (BFS ensures minimal presses)
		for _, b := range validButtons {
			// pressing is only allowed if all affected indices can be incremented
			can := true
			for _, idx := range b {
				if cur.state[idx] >= targets[idx] {
					// pressing would exceed target for this counter -> invalid
					can = false
					break
				}
			}
			if !can {
				continue
			}
			// form next state
			next := make([]int, n)
			copy(next, cur.state)
			for _, idx := range b {
				next[idx]++
			}
			key := encodeState(next)
			if _, seen := visited[key]; !seen {
				visited[key] = struct{}{}
				q = append(q, node{state: next, dist: cur.dist + 1})
			}
		}
	}

	// If BFS exhausted, unreachable
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	lineNo := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lineNo++
		if line == "" {
			continue
		}
		buttons, targets, err := parseLine(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse error on line %d: %v\n", lineNo, err)
			os.Exit(1)
		}
		presses := minPressesForMachine(buttons, targets)
		if presses < 0 {
			fmt.Fprintf(os.Stderr, "machine on line %d is unreachable (targets: %v)\n", lineNo, targets)
			os.Exit(1)
		}
		fmt.Printf("Machine %d requires minimum %d presses.\n", lineNo, presses)
		total += presses
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
		os.Exit(1)
	}
	fmt.Printf("Total minimum presses: %d\n", total)
}
