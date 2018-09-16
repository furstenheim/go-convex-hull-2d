package go_convex_hull_2d

import (
	"fmt"
	"math/rand"
	"testing"
	"github.com/stretchr/testify/assert"
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


func TestSortByIndexes (t *testing.T) {
	testCases := []struct{
		description string
		original []int
		indexes []int
		expected []int
	}{
		{
			"",
			[]int{0, 1, 1, 2},
			[]int{0, 1},
			[]int{0, 1},
		},
		{
			"Indexes start at middle",
			[]int{0, 1, 1, 2, 4, 5, 6},
			[]int{2, 3},
			[]int{1, 2},
		},
		{
			"Reversed indexes",
			[]int{0, 1, 2, 3, 4, 5, 6},
			[]int{3, 2},
			[]int{3, 2},
		},
		{
			"Simple swap would fail",
			[]int{0, 1, 2, 3, 4, 5, 6},
			[]int{1, 0, 2},
			[]int{1, 0, 2},
		},
	}
	for _, tc := range(testCases) {
		t.Run(tc.description, func (t *testing.T) {
			original := Interface(sortByIndexInterface(tc.original))
			final := []int(sortByIndexes(original, tc.indexes).(sortByIndexInterface))
			assert.Equal(t, tc.expected, final)
		})
	}
}

func BenchmarkConvexHull(b *testing.B) {
	testCases := []struct{
		size int
	}{
		{100},
		{1000},
		{10000},
		{100000},
		{1000000},
	}

	for _, tc := range(testCases) {
		points := make([]float64, 2 * tc.size)
		for i, _ := range(points) {
			points[i] = rand.Float64()
		}
		b.Run(fmt.Sprintf("Convex hull of size %d", tc.size), func (b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = New(FlatPoints(points)).(FlatPoints)
			}
		})
	}
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
			t.Errorf("%d th point of the convex hull was not correct, got: %+v, %+v want: %+v %+v", i, x1, y1, x2, y2)
		}
	}
}


type point struct {
	x, y float64
}

func (p point) GetCoordinates() (x, y float64) {
	return p.x, p.y
}

type sortByIndexInterface []int

func (c sortByIndexInterface) Take(i int) (x, y float64) {
	return 0, 0
}

func (c sortByIndexInterface) Len() int {
	return len(c)
}

func (c sortByIndexInterface) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c sortByIndexInterface) Slice(i, j int) Interface {
	return c[i:j]
}
