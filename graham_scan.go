package main

import (
	"fmt"
	"math"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Point struct {
	X, Y float64
}

type PointList []Point

func makePoint(x float64, y float64) Point {
	return Point{X: x, Y: y}
}

func printStack(s *Stack) {
	v := s.top
	fmt.Printf("Stack: ")
	for v != nil {
		fmt.Printf("%v ", v.value)
		v = v.next
	}
	fmt.Println("")
}

// Implement sort interface
func (p PointList) Len() int {
	return len(p)
}

func (p PointList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PointList) Less(i, j int) bool {
	area := Area2(p[0], p[i], p[j])
	if area == 0 {
		x := math.Abs(p[i].X-p[0].X) - math.Abs(p[j].X-p[0].X)
		y := math.Abs(p[i].Y-p[0].Y) - math.Abs(p[j].Y-p[0].Y)
		if x < 0 || y < 0 {
			return true
		} else if x > 0 || y > 0 {
			return false
		} else {
			return false
		}
	}
	return area > 0
}

func (p PointList) FindLowestPoint() {
	m := 0
	for i := 1; i < len(p); i++ {
		//If lowest points are on the same line, take the rightmost point
		if (p[i].Y < p[m].Y) || ((p[i].Y == p[m].Y) && p[i].X > p[m].X) {
			m = i
		}
	}
	p[0], p[m] = p[m], p[0]
}

func (points PointList) Compute() (PointList, bool) {
	if len(points) < 3 {
		return nil, false
	}

	stack := new(Stack)
	points.FindLowestPoint()
	sort.Sort(&points)

	stack.Push(points[0])
	stack.Push(points[1])

	i := 2
	for i < len(points) {
		pi := points[i]
		p1 := stack.top.next.value.(Point)
		p2 := stack.top.value.(Point)
		if isLeft(p1, p2, pi) {
			stack.Push(pi)
			i++
		} else {
			stack.Pop()
		}
	}

	//Copy the hull
	ret := make(PointList, stack.Len())
	top := stack.top
	count := 0
	for top != nil {
		ret[count] = top.value.(Point)
		top = top.next
		count++
	}
	return ret, true
}

func isLeft(p0, p1, p2 Point) bool {
	return Area2(p0, p1, p2) > 0
}

func Area2(a, b, c Point) float64 {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}

func (points PointList) DrawPoints() {
	for _, p := range points {
		// Convert from normalized coordinates back to screen coordinates
		screenX := int32((p.X + 1) * float64(width) / 2)
		screenY := int32((-p.Y + 1) * float64(height) / 2)
		rl.DrawCircle(
			screenX,
			screenY,
			5,
			rl.Red,
		)
	}
}

func (points PointList) DrawLines() {
	for i := 0; i < len(points); i++ {

		next := (i + 1) % len(points)
		// Convert from normalized coordinates back to screen coordinates
		screenX1 := int32((points[i].X + 1) * float64(width) / 2)
		screenY1 := int32((-points[i].Y + 1) * float64(height) / 2)
		screenX2 := int32((points[next].X + 1) * float64(width) / 2)
		screenY2 := int32((-points[next].Y + 1) * float64(height) / 2)

		rl.DrawLine(
			screenX1,
			screenY1,
			screenX2,
			screenY2,
			rl.Blue,
		)
	}
}

func (points PointList) DrawLowestPoint() {
	if len(points) <= 0 {
		return
	}
	// Convert from normalized coordinates back to screen coordinates
	screenX := int32((points[0].X + 1) * float64(width) / 2)
	screenY := int32((-points[0].Y + 1) * float64(height) / 2)
	rl.DrawCircle(
		screenX,
		screenY,
		5,
		rl.Black,
	)
}
