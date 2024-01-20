/*Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором. */
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния до другой точки(формула Евклида)
func (p *Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	point1 := NewPoint(2, 7)
	point2 := NewPoint(10, 14)

	distance := point1.DistanceTo(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
