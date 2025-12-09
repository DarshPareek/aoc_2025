package main

import (
	"bufio"
	"fmt"
	//"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Edge struct {
	p1, p2 Point
}

func readInput(filepath string) ([]Point, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		points = append(points, Point{x, y})
	}
	return points, scanner.Err()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// isInsidePolygon checks if a specific coordinate (represented as float to test centers)
// is strictly inside the polygon using the Ray Casting algorithm.
func isPointInsidePolygon(testX, testY float64, edges []Edge) bool {
	inside := false
	for _, edge := range edges {
		p1x, p1y := float64(edge.p1.x), float64(edge.p1.y)
		p2x, p2y := float64(edge.p2.x), float64(edge.p2.y)

		// Check if ray crosses the edge
		if ((p1y > testY) != (p2y > testY)) &&
			(testX < (p2x-p1x)*(testY-p1y)/(p2y-p1y)+p1x) {
			inside = !inside
		}
	}
	return inside
}

// hasEdgeCollision checks if any polygon edge strictly cuts through the candidate rectangle.
// It allows edges to overlap with the rectangle border, but not pass through the middle.
func hasEdgeCollision(x1, y1, x2, y2 int, edges []Edge) bool {
	minX, maxX := min(x1, x2), max(x1, x2)
	minY, maxY := min(y1, y2), max(y1, y2)

	for _, edge := range edges {
		// 1. Check if a polygon Vertex is strictly inside the rectangle
		// If a vertex is inside, the polygon folds into our rectangle -> Invalid.
		if edge.p1.x > minX && edge.p1.x < maxX && edge.p1.y > minY && edge.p1.y < maxY {
			return true
		}

		// 2. Check for "Cross" Intersection
		// Vertical Edge crossing Horizontal Rect span
		isVert := edge.p1.x == edge.p2.x
		if isVert {
			ex := edge.p1.x
			ey1, ey2 := min(edge.p1.y, edge.p2.y), max(edge.p1.y, edge.p2.y)
			// Does the edge X fall strictly between Rect X bounds?
			if ex > minX && ex < maxX {
				// Does the edge Y range completely span across the Rect Y bounds?
				// (i.e. splitting the rectangle in half)
				if ey1 <= minY && ey2 >= maxY {
					return true
				}
				// Or does the edge partially enter the rectangle?
				// (This is actually covered by the "Vertex Inside" check above,
				// but strictly crossing edges might not have vertices inside if the edge is long)
				if (ey1 > minY && ey1 < maxY) || (ey2 > minY && ey2 < maxY) {
					return true
				}
			}
		} else {
			// Horizontal Edge crossing Vertical Rect span
			ey := edge.p1.y
			ex1, ex2 := min(edge.p1.x, edge.p2.x), max(edge.p1.x, edge.p2.x)
			if ey > minY && ey < maxY {
				if ex1 <= minX && ex2 >= maxX {
					return true
				}
				if (ex1 > minX && ex1 < maxX) || (ex2 > minX && ex2 < maxX) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	points, err := readInput("./input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(points) < 4 {
		fmt.Println("Not enough points to form a polygon.")
		return
	}

	// 1. Build Edges List
	var edges []Edge
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]
		edges = append(edges, Edge{p1, p2})
	}

	maxArea := 0

	// 2. Iterate all pairs of points (Potential Rectangle Corners)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			// Skip if they form a line, not a rectangle
			if p1.x == p2.x || p1.y == p2.y {
				continue
			}

			// Calculate Area
			// The problem asks for "tiles", so it's inclusive area
			// width = difference + 1
			w := abs(p1.x-p2.x) + 1
			h := abs(p1.y-p2.y) + 1
			area := w * h

			// Optimization: Don't check validity if this area is smaller than what we already found
			if area <= maxArea {
				continue
			}

			// 3. Validation: Is this rectangle valid?
			// A rectangle is valid if:
			// A. Its center point is inside the polygon (ensures we aren't selecting empty void)
			// B. No polygon edges cut through it.

			midX := float64(p1.x+p2.x) / 2.0
			midY := float64(p1.y+p2.y) / 2.0

			if !isPointInsidePolygon(midX, midY, edges) {
				continue
			}

			if hasEdgeCollision(p1.x, p1.y, p2.x, p2.y, edges) {
				continue
			}

			// If we passed checks, this is a valid rectangle
			maxArea = area
		}
	}

	fmt.Printf("Largest Area: %d\n", maxArea)
}
