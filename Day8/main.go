package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	id int
	x  int
	y  int
	z  int
}

type Edge struct {
	u    int
	v    int
	dist float64
}
type UnionFind struct {
	parent []int
	size   []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return uf.parent[i]
}
func (uf *UnionFind) Union(i, j int) bool {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)

	if rootI != rootJ {
		if uf.size[rootI] < uf.size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		uf.parent[rootJ] = rootI
		uf.size[rootI] += uf.size[rootJ]
		uf.count--
		return true
	}
	return false
}

func readInput(filepath string) ([]Point, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []Point
	scanner := bufio.NewScanner(file)
	idCounter := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		points = append(points, Point{id: idCounter, x: x, y: y, z: z})
		idCounter++
	}
	return points, scanner.Err()
}

func main() {
	points, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}
	var edges []Edge
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := points[i]
			p2 := points[j]
			dx := float64(p1.x - p2.x)
			dy := float64(p1.y - p2.y)
			dz := float64(p1.z - p2.z)
			dist := math.Sqrt(dx*dx + dy*dy + dz*dz)
			edges = append(edges, Edge{u: i, v: j, dist: dist})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})
	uf := NewUnionFind(n)
	for _, edge := range edges {
		merged := uf.Union(edge.u, edge.v)
		if merged {
			if uf.count == 1 {
				p1 := points[edge.u]
				p2 := points[edge.v]
				fmt.Printf("Last connection made between: Index %d and Index %d\n", edge.u, edge.v)
				fmt.Printf("Coordinates: (%d,%d,%d) and (%d,%d,%d)\n", p1.x, p1.y, p1.z, p2.x, p2.y, p2.z)
				result := p1.x * p2.x
				fmt.Printf("Result (X1 * X2): %d\n", result)
				return
			}
		}
	}
}
