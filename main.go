// Package provides a function to compute the convex hull using the monotone chain algorithm
// ported to go from https://github.com/mikolalysenko/monotone-convex-hull-2d
package go_convex_hull_2d

import (
	"sort"
	"sync"
)

type Point interface {
	getCoordinates() (float64, float64)
}

// Given an array of Points it computes the convex hull
func ComputeConvexHull(points []Point) []Point {
	sort.Sort(pointSorter(points))
	return ComputeConvexHullOnSortedArray(points)
}

// Given an array of Points ordered lexicographically by (x,y) it computes the convex hull
func ComputeConvexHullOnSortedArray(points []Point) []Point {
	n := len(points)
	if n < 3 {
		return points
	}
	var w sync.WaitGroup
	// Run lower and upper parts in parallel
	var lower = []Point{points[0], points[1]}
	var upper = []Point{points[len(points)- 1], points[len(points)-2]}
	w.Add(2)
	// lower part
	go func() {
		for _, p := range points[2:] {
			m := len(lower)
			for m > 1 && !isOrientationPositive(lower[m-2], lower[m-1], p) {
				lower = lower[:m-1]
				m -= 1
			}
			lower = append(lower, p)
		}

		w.Done()
	}()
	// upper part
	go func() {
		for i := len(points) - 3; i >= 0; i-- {
			p := points[i]
			m := len(upper)
			for m > 1 && !isOrientationPositive(upper[m-2], upper[m-1], p) {
				upper = upper[:m-1]
				m -= 1
			}
			upper = append(upper, p)

		}
		w.Done()
	}()
	w.Wait()

	// End points are duplicated
	upper = upper[:len(upper)-1]
	lower = lower[:len(lower)-1]
	return append(lower, upper...)
}

func isOrientationPositive(p1, p2, p3 Point) (isPositive bool) {
	x1, y1 := p1.getCoordinates()
	x2, y2 := p2.getCoordinates()
	x3, y3 := p3.getCoordinates()
	// compute determinant to obtain the orientation
	// |x1 - x3 x2 - x3 |
	// |y1 - y3 y2 - y3 |
	return (x1-x3)*(y2-y3)-(y1-y3)*(x2-x3) > 0
}

type pointSorter []Point

func (s pointSorter) Len() int {
	return len(s)
}

func (s pointSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s pointSorter) Less(i, j int) bool {
	x1, y1 := s[i].getCoordinates()
	x2, y2 := s[j].getCoordinates()
	if x1 < x2 {
		return true
	}
	if x1 == x2 {
		return y1 < y2
	}
	return false
}
