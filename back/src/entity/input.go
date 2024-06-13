package entity

import "github.com/AlexSilverson/Dojdik/src/geom"

type Input struct {
	N      int          `json:"n"`
	H      float64      `json:"h"`
	Points []geom.Point `json:"points"`
}
