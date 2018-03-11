package go_convex_hull_2d

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestConvexHull(t *testing.T) {
	points := FlatPoints([]float64{0, 0, 1, 1, 1, 0, 0.5, 0.5, 0.7, 0.1})
	convexHull := New(points).(FlatPoints)
	compareConvexHulls(t, convexHull, FlatPoints([]float64{0, 0, 1, 0, 1, 1}))

	points = FlatPoints([]float64{0, 0, 1, 0, 1, 1, 0, 1})
	convexHull = New(points).(FlatPoints)
	compareConvexHulls(t, convexHull, FlatPoints([]float64{0, 0, 1, 0, 1, 1, 0, 1}))

	for i := 0; i < 1000; i++ {
		points = append(points, rand.Float64(), rand.Float64())
	}
	convexHull = New(points).(FlatPoints)
	compareConvexHulls(t, convexHull, FlatPoints([]float64{0, 0, 1, 0, 1, 1, 0, 1}))

	// TODO degenerate cases

}

func compareConvexHulls(t *testing.T, actualC, expectedC FlatPoints) {
	if actualC.Len() != expectedC.Len() {
		t.Errorf("Convex hull didn't correct length, got %d, want: %d", len(actualC), len(expectedC))
		for i := 0; i < actualC.Len(); i++ {
			t.Log(actualC.Take(i))
		}
		return
	}
	for i := 0; i < actualC.Len(); i++ {
		x1, y1 := actualC.Take(i)
		x2, y2 := expectedC.Take(i)
		if x1 != x2 || y1 != y2 {
			fmt.Println(actualC, expectedC)
			t.Errorf("%d th point of the convex hull was not correct, got: %+v want: %+v", i, x1, y1, x2, y2)
		}
	}
}

type point struct {
	x, y float64
}

func (p point) GetCoordinates() (x, y float64) {
	return p.x, p.y
}
