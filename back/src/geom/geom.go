package geom

import "math"

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Line struct {
	A, B, C float64
}

func MakeLine(a, b Point) Line {
	var ans Line
	ans.A = (b.Y - a.Y) * -1
	ans.B = b.X - a.X
	ans.C = (ans.A*a.X + ans.B*a.Y) * -1
	return ans
}

func Cross(a, b Point) float64 {
	return a.X*b.Y - b.X*a.Y
}

func GetSerf(a []Point) (ans float64) {
	ans = 0
	n := len(a)
	for i := 0; i < n-1; i++ {
		var v1, v2 Point
		v1.X = a[i].X - a[0].X
		v1.Y = a[i].Y - a[0].Y
		v2.X = a[i+1].X - a[0].X
		v2.Y = a[i+1].Y - a[0].Y
		ans += Cross(v1, v2) / 2
	}
	return math.Abs(ans)
}

func LineCross(a, b Line) Point {
	var ans Point
	ans.X = (a.B*b.C - b.B*a.C) / (a.A*b.B - b.A*a.B)
	ans.Y = (a.C*b.A - b.C*a.A) / (a.A*b.B - a.B*b.A)
	return ans
}
