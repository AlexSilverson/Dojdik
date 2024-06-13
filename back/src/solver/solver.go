package solver

import (
	"fmt"
	"math"
	"slices"

	"github.com/AlexSilverson/Dojdik/src/entity"
	"github.com/AlexSilverson/Dojdik/src/geom"
)

func Solve(input entity.Input) float64 {
	n := input.N
	h := input.H
	n++
	a := input.Points
	var mnx, mxx float64 = 10001, -10001
	for i := 0; i < n; i++ {
		mnx = math.Min(mnx, a[i].X)
		mxx = math.Max(mxx, a[i].X)
	}

	a = append([]geom.Point{geom.Point{mnx, 20000000000}}, a...)
	a = append(a, geom.Point{mxx, 20000000000})
	var ans float64
	//Предподсчет количества воды стекающего в точку слева
	lstock := make([]float64, n+2)
	rstock := make([]float64, n+2)
	lstock[0] = 0
	rstock[0] = 0

	var now []geom.Point
	now = append(now, a[0])
	for i := 1; i <= n+1; i++ {

		if a[i-1].Y > a[i].Y {
			lstock[i] = lstock[i-1] + ((a[i].X - a[i-1].X) * h)

		} else {
			var nowres []geom.Point

			for j := len(now) - 1; j > 0; j-- {

				nowres = append(nowres, now[j])
				now = now[:len(now)-1]
				if now[j-1].Y > a[i].Y {

					slices.Reverse(nowres)
					nowh := geom.MakeLine(geom.Point{0, a[i].Y}, geom.Point{1, a[i].Y})
					lline := geom.MakeLine(nowres[0], now[j-1])

					nowres = append([]geom.Point{geom.LineCross(nowh, lline)}, nowres...)
					nowres = append(nowres, a[i])

					serf := geom.GetSerf(nowres)

					lstock[i] = math.Max(0, lstock[i-1]-serf+((a[i].X-a[i-1].X)*h))
					now = append(now, geom.LineCross(nowh, lline))
					break
				}
			}
		}
		now = append(now, a[i])
	}

	//Предподсчет количества воды стекающего в точку справа
	now = nil
	slices.Reverse(a)
	now = append(now, a[0])

	for i := 1; i <= n+1; i++ {

		if a[i-1].Y > a[i].Y {
			rstock[i] = rstock[i-1] + ((a[i-1].X - a[i].X) * h)

		} else {
			var nowres []geom.Point

			for j := len(now) - 1; j > 0; j-- {

				nowres = append(nowres, now[j])
				now = now[:len(now)-1]
				if now[j-1].Y > a[i].Y {

					slices.Reverse(nowres)
					nowh := geom.MakeLine(geom.Point{0, a[i].Y}, geom.Point{1, a[i].Y})
					lline := geom.MakeLine(nowres[0], now[j-1])

					nowres = append([]geom.Point{geom.LineCross(nowh, lline)}, nowres...)
					nowres = append(nowres, a[i])

					serf := geom.GetSerf(nowres)

					rstock[i] = math.Max(0, rstock[i-1]-serf+((a[i-1].X-a[i].X)*h))
					now = append(now, geom.LineCross(nowh, lline))
					break
				}
			}
		}
		now = append(now, a[i])
	}
	//Основной алгоритм
	now = nil
	slices.Reverse(rstock)
	slices.Reverse(a)

	var eps float64 = 0.000001
	for i := 0; i <= n; i++ {
		nowmxy, nowmny := a[i].Y, a[i].Y
		now = nil
		now = append(now, a[i])

		for j := i + 1; j <= n+1; j++ {
			if a[j].Y > nowmxy {
				l, r := nowmxy, math.Min(a[j].Y, a[i-1].Y)

				for r-l > eps {
					m := (l + r) / 2
					lline := geom.MakeLine(a[i], a[i-1])
					rline := geom.MakeLine(a[j], a[j-1])
					nowh := geom.MakeLine(geom.Point{0, m}, geom.Point{1, m})

					now = append([]geom.Point{geom.LineCross(lline, nowh)}, now...)
					now = append(now, geom.LineCross(rline, nowh))

					serf := geom.GetSerf(now)
					now = now[1 : len(now)-1]

					if serf < h*(a[j].X-a[i-1].X)+lstock[i-1]+rstock[j] {
						l = m
					} else {
						r = m
					}
				}

				if l-eps < nowmxy {

				} else {
					ans = math.Max(ans, l-nowmny)
				}

			}

			if i == 0 || a[j].Y > a[i-1].Y {
				break
			}
			nowmny = math.Min(nowmny, a[j].Y)
			nowmxy = math.Max(nowmxy, a[j].Y)
			now = append(now, a[j])
		}
	}
	fmt.Println(now)
	return ans
}
