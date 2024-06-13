package DTO

import (
	"github.com/AlexSilverson/Dojdik/src/entity"
	"github.com/AlexSilverson/Dojdik/src/geom"
)

type InputDTO struct {
	H  float64   `json:"h"`
	AR []float64 `json:"arr"`
}

func (t InputDTO) MapToInput() (inp entity.Input) {
	inp.H = t.H
	inp.N = len(t.AR)/2 - 1
	inp.Points = make([]geom.Point, inp.N+1)
	for i := 0; i < len(t.AR); i += 2 {
		inp.Points[i/2] = geom.Point{X: t.AR[i], Y: t.AR[i+1]}
	}
	return
}

/*
-5 10
-3 4
-1 6
1 -4
4 17
5 3
9 5
12 15
*/
