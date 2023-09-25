package main

import (
	"math"

	"github.com/lewiscasewell/tolling/types"
)

type CalculatorServicer interface {
	CalculateDistance(data types.OBUData) (float64, error)
}

type CalculatorService struct {
	prevPoint Point
}

type Point struct {
	Latitude  float64
	Longitude float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0

	if distance == 0.0 && s.prevPoint.Latitude == 0.0 && s.prevPoint.Longitude == 0.0 {
		s.prevPoint = Point{
			Latitude:  data.Latitude,
			Longitude: data.Longitude,
		}
	} else {
		distance = calculateDistance(s.prevPoint.Latitude, s.prevPoint.Longitude, data.Latitude, data.Longitude)
		s.prevPoint = Point{
			Latitude:  data.Latitude,
			Longitude: data.Longitude,
		}
	}

	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
