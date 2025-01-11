package main

import (
	"math"
	"traffic-toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
	prevPoint types.Point
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := calculatorService(s.prevPoint.Lat, s.prevPoint.Lng, data.Lat, data.Lng)

	s.prevPoint.Lat = data.Lat
	s.prevPoint.Lng = data.Lng

	return distance, nil
}

func calculatorService(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
