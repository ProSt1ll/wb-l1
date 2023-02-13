package main

import (
	"fmt"
	"math"
)

func main() {
	a := New(1, 0)
	b := New(4, 1)
	fmt.Println(Distance(a, b))
}

//структура точки
type Point struct {
	x int
	y int
}

//конструктор
func New(x int, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

//функция вычисления дистанции
func Distance(a Point, b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}
