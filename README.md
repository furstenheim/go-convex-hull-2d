## Go convex hull 2d

Implementation of convex hull using monotone chain algorithm. The order is `O(nlog(n))` if points are not sorted and `O(n)` if they are.
The algorithm works in place and returns a slice with the convex hull. Algorithm works on any object implementing go_convex_hull_2d.Interface,
 which abstract an array of points. If the elements of your slice already implement the interface Points then you can use Convexer
## Example

        import {
            "go_convex_hull_2d"
        }

        type point struct {
        	x, y float64
        }

        func (p point) GetCoordinates () (x, y float64) {
        	return p.x, p.y
        }

        points := []Point{point{0, 0}, point{0, 1}, point{1,1}, point{1/2, 1/2}}
        convexHull := ComputeConvexHull(Convexer(points))
        > // {{0,0}, {0, 1}, {1,1}

### Acknowledgment

Source code is based on https://github.com/mikolalysenko/monotone-convex-hull-2d, originally written in JS.
