package go_convex_hull_2d

import "fmt"

func ExampleComputeConvexHull_array_of_points() {
	points := []float64{0, 0, 0, 1, 1, 1, 1 / 2, 1 / 2}
	convexHull := New(FlatPoints(points))
	fmt.Println(convexHull)
	// Output: [0 0 1 1 0 1]
}

type FlatPoints []float64

func (fp FlatPoints) Len () int {
	return len(fp) / 2
}

func (fp FlatPoints) Swap (i, j int) {
	fp[2 * i], fp[2 * i + 1], fp[2 * j], fp[2 * j + 1] = fp[2 * j], fp[2 * j + 1], fp[2 * i], fp[2 * i + 1]
}

func (fp FlatPoints) Take(i int) (x1, y1 float64) {
	return fp[2 * i], fp[2 * i +1]
}
func (fp FlatPoints) Slice(i, j int) Interface {
	return fp[2 * i: 2 * j]
}

