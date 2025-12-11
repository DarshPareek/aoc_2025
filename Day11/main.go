package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type CacheKey struct {
	node string
	dac  bool
	fft  bool
}

type GraphTraverser struct {
	graph map[string][]string
	memo1 map[string]int
	memo2 map[CacheKey]int
	mu    sync.RWMutex
}

func NewTraverser(graph map[string][]string) *GraphTraverser {
	return &GraphTraverser{
		graph: graph,
		memo1: make(map[string]int),
		memo2: make(map[CacheKey]int),
	}
}

func (t *GraphTraverser) travel(origin string) int {
	if origin == "out" {
		return 1
	}

	t.mu.RLock()
	if val, ok := t.memo1[origin]; ok {
		t.mu.RUnlock()
		return val
	}
	t.mu.RUnlock()

	res := 0
	children := t.graph[origin]
	for _, child := range children {
		res += t.travel(child)
	}

	t.mu.Lock()
	t.memo1[origin] = res
	t.mu.Unlock()

	return res
}

func (t *GraphTraverser) travel2(origin string, sdac bool, sfft bool) int {
	key := CacheKey{node: origin, dac: sdac, fft: sfft}

	if origin == "out" {
		if sdac && sfft {
			return 1
		}
		return 0
	}

	t.mu.RLock()
	if val, ok := t.memo2[key]; ok {
		t.mu.RUnlock()
		return val
	}
	t.mu.RUnlock()

	res := 0
	children := t.graph[origin]

	for _, child := range children {
		isFft := sfft || child == "fft"
		isDac := sdac || child == "dac"

		res += t.travel2(child, isDac, isFft)
	}

	t.mu.Lock()
	t.memo2[key] = res
	t.mu.Unlock()

	return res
}

func readInput(filepath string) (map[string][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graph := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		neighbors := strings.Fields(parts[1])
		graph[key] = neighbors
	}

	return graph, scanner.Err()
}

func main() {
	graph, err := readInput("./input.txt")
	if err != nil {
		panic(err)
	}

	traverser := NewTraverser(graph)

	var wg sync.WaitGroup
	var res1, res2 int

	wg.Add(1)
	go func() {
		defer wg.Done()
		res1 = traverser.travel("you")
		fmt.Println("Travel 1 Finished")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res2 = traverser.travel2("svr", false, false)
		fmt.Println("Travel 2 Finished")
	}()

	wg.Wait()

	fmt.Printf("Graph: %v\nResult 1: %d\nResult 2: %d\n", len(graph), res1, res2)
}
