// Package go_convex_hull_2d provides convex hull computation for any type that implements
// go_convex_hull_2d.Interface. A convex hull is the smallest convex hull covering a set of points
// based originally on https://github.com/mikolalysenko/monotone-convex-hull-2d
// convexHull works in place in the interface given reordering the points and removing those that do not belong
// to the convexhull
package go_convex_hull_2d

import (
	"log"
	"sort"
	"sync"
)

// Interface abstracting the necessary methods of a point array
type Interface interface {
	Take(i int) (x, y float64)         // Retrieve point at position i
	Len() int                 // Number of elements
	Swap(i, j int)            // Swap elements with indexes i and j
	Slice(i, j int) Interface //Slice the interface between two indices
}


// Given an Interface computes the convex hull
func New(points Interface) Interface {
	sort.Sort(pointSorter{i: points})
	return NewFromSortedArray(points)
}

// Given an Interface which is already ordered in lexicographical order by (x,y) it computes the convex hull
func NewFromSortedArray(points Interface) Interface {
	n := points.Len()
	if n < 3 {
		return points
	}
	var w sync.WaitGroup
	// Run lower and upper parts in parallel. Compute array of indices
	var lowerIndexes = []int{0, 1}
	var upperIndexes = []int{n - 1, n - 2}
	w.Add(2)
	// lower part
	func() {
		for i := 2; i < n; i++ {
			x, y := points.Take(i)
			m := len(lowerIndexes)
			for m > 1 {
				x2, y2 := points.Take(lowerIndexes[m-2])
				x3, y3 := points.Take(lowerIndexes[m-1])
				if isOrientationPositive(x2, y2, x3, y3, x, y) {
					break
				}
				lowerIndexes = lowerIndexes[:m-1]
				m -= 1
			}
			lowerIndexes = append(lowerIndexes, i)
		}

		w.Done()
	}()
	// upper part
	func() {
		for i := n - 3; i >= 0; i-- {
			x, y := points.Take(i)
			m := len(upperIndexes)
			for m > 1 {
				x2, y2 := points.Take(upperIndexes[m-2])
				x3, y3 := points.Take(upperIndexes[m-1])
				if isOrientationPositive(x2, y2, x3, y3, x, y) {
					break
				}
				upperIndexes = upperIndexes[:m-1]
				m -= 1
			}
			upperIndexes = append(upperIndexes, i)

		}
		w.Done()
	}()
	w.Wait()

	// End points are duplicated
	upperIndexes = upperIndexes[:len(upperIndexes)-1]
	lowerIndexes = lowerIndexes[:len(lowerIndexes)-1]
	allIndexes := append(lowerIndexes, upperIndexes...)
	return sortByIndexes(points, allIndexes)
}

func isOrientationPositive(x1, y1, x2, y2, x3, y3 float64) (isPositive bool) {
	// compute determinant to obtain the orientation
	// |x1 - x3 x2 - x3 |
	// |y1 - y3 y2 - y3 |
	return (x1-x3)*(y2-y3)-(y1-y3)*(x2-x3) > 0
}

type pointSorter struct {
	i Interface
}

func (s pointSorter) Less(i, j int) bool {
	x1, y1 := s.i.Take(i)
	x2, y2 := s.i.Take(j)
	return x1 < x2 || (x1 == x2 && y1 < y2)
}

func (s pointSorter) Swap(i, j int) {
	s.i.Swap(i, j)
}

func (s pointSorter) Len() int {
	return s.i.Len()
}

type byMap struct {
	i Interface
	m map[int]int
}

func (o byMap) Len() int {
	return o.i.Len()
}

func (o byMap) Less(i, j int) bool {
	i1, ok1 := o.m[i]
	i2, ok2 := o.m[j]
	if !ok1 || !ok2 {
		log.Fatal("Unkown state, one index was not tracked in the map")
	}
	return i1 < i2
}

// When swapping elements we must update the map
func (o byMap) Swap(i, j int) {
	i1, ok1 := o.m[i]
	i2, ok2 := o.m[j]
	if !ok1 || !ok2 {
		log.Fatal("Unkown state, one index was not tracked in the map")
	}
	// swap priorities
	o.m[j] = i1
	o.m[i] = i2
	// swap containing slice
	o.i.Swap(i, j)
}

func sortByIndexes (points Interface, indexes []int) Interface {
	n := points.Len()
	// Now sort Interface leaving first the indexes we are interested in
	var orderMap = make(map[int]int, n)
	for i, j := range indexes {
		orderMap[j] = i
	}
	// mark all other points as bigger
	for i := 0; i < n; i++ {
		_, ok := orderMap[i]
		if !ok {
			orderMap[i] = len(indexes)
		}
	}
	bM := byMap{i: points, m: orderMap}
	sort.Sort(bM)
	return points.Slice(0, len(indexes))
}
