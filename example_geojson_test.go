package go_convex_hull_2d

import "fmt"

type coordinates [][2]float64

func (c coordinates) Take(i int) Point {
	return point{c[i][0], c[i][1]}
}

func (c coordinates) Len() int {
	return len(c)
}

func (c coordinates) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c coordinates) Slice(i, j int) Interface {
	return c[i:j]
}

func ExampleComputeConvexHull_coordinates() {
	points := coordinates{{0, 0}, {0, 1}, {1, 1}, {1 / 2, 1 / 2}}
	convexHull := ComputeConvexHull(points)
	fmt.Println(convexHull)
	// Output: [[0 0] [1 1] [0 1]]
}
