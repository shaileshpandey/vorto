package models

import (
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y float64
}

func (p *Coordinate) formatPoint(coordinateText string) {
	coordinateText = strings.Trim(coordinateText, "()")
	coordinates := strings.Split(coordinateText, ",")
	x, _ := strconv.ParseFloat(coordinates[0], 64)
	y, _ := strconv.ParseFloat(coordinates[1], 64)
	p.X = x
	p.Y = y
}

func (point1 Coordinate) euclidDistance(point2 Coordinate) float64 {
	return math.Sqrt(math.Pow(point2.X-point1.X, 2) + math.Pow(point2.Y-point1.Y, 2))
}
