package go_convex_hull_2d

import (
	"testing"
	"fmt"
	"math/rand"
)

func TestConvexHull (t *testing.T) {
	points := toPoints([]float64{0,0,1,1,1,0,0.5,0.5,0.7,0.1})
	convexHull := ComputeConvexHull(points).(Convexer)
	compareConvexHulls(t, convexHull, toPoints([]float64{0, 0, 1, 0, 1, 1}))

	points = toPoints([]float64{0,0, 1,0, 1,1, 0,1})
	convexHull = ComputeConvexHull(points).(Convexer)
	compareConvexHulls(t, convexHull, toPoints([]float64{0,0, 1,0, 1,1, 0, 1}))

	for i:= 0; i < 1000; i++ {
		points = append(points, point{rand.Float64(), rand.Float64()})
	}
	convexHull = ComputeConvexHull(points).(Convexer)
	compareConvexHulls(t, convexHull, toPoints([]float64{0,0, 1,0, 1,1, 0, 1}))

	// TODO degenerate cases

}

func compareConvexHulls (t *testing.T, actualC, expectedC Convexer) {
	if (len(actualC) != len(expectedC)) {
		t.Errorf("Convex hull didn't correct length, got %d, want: %d", len(actualC), len(expectedC))
		for _, p := range(actualC) {
			t.Log(p.(point))
		}
		return
	}
	for i, p1 := range(actualC) {
		p2 := expectedC[i]
		x1, y1 := p1.GetCoordinates()
		x2, y2 := p2.GetCoordinates()
		if ( x1 != x2 || y1 != y2) {
			fmt.Println(actualC, expectedC)
			t.Errorf("%d th point of the convex hull was not correct, got: %+v want: %+v", i, p1, p2)
		}
	}
}

type point struct {
	x, y float64
}

func (p point) GetCoordinates () (x, y float64) {
	return p.x, p.y
}

func toPoints (ps []float64) (o Convexer) {
	o = make([]Point, len(ps) / 2)
	for i := 0; i < len(ps); i+= 2{
		o[i/2] = point{ps[i], ps[i+1]}
	}
	return Convexer(o)
}

