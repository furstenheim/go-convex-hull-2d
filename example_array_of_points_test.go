package go_convex_hull_2d

import "fmt"


func ExampleComputeConvexHull_array_of_points() {
	points := []Point{point{0, 0}, point{0, 1}, point{1,1}, point{1/2, 1/2}}
	convexHull := ComputeConvexHull(Convexer(points))
	fmt.Println(convexHull)
	// Output: [{0 0} {1 1} {0 1}]
}
