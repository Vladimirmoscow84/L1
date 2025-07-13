/*Разработать программу нахождения расстояния между двумя точками на плоскости. Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором. Расстояние рассчитывается по формуле между координатами двух точек.

Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// функция-конструктор координаты точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// метод, возвращающий дистанцию между точками
func (p Point) Distance(other Point) float64 {
	distx := p.x - other.x
	disty := p.y - other.y
	distance := math.Sqrt(math.Pow(distx, 2) + math.Pow(disty, 2))
	return distance

}

func main() {
	fmt.Println("задайте координаты первой точки")
	var x1, y1 float64
	fmt.Scan(&x1, &y1)
	point1 := NewPoint(x1, y1)

	fmt.Println("задайте координаты второй точки")
	var x2, y2 float64
	fmt.Scan(&x2, &y2)
	point2 := NewPoint(x2, y2)

	distance := point1.Distance(*point2)

	fmt.Printf("расстояние между точками: %.2f\n", distance)

}
