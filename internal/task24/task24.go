package task24

import (
	"fmt"
	"math"
)

func Task24() {
	var x1, y1 float64
	fmt.Print("Введите координату x для точки 1: ")
	fmt.Scan(&x1)
	fmt.Print("Введите координату y для точки 1: ")
	fmt.Scan(&y1)

	var x2, y2 float64
	fmt.Print("Введите координату x для точки 2: ")
	fmt.Scan(&x2)
	fmt.Print("Введите координату y для точки 2: ")
	fmt.Scan(&y2)

	// Создаем две точки
	point1 := NewPoint(x1, y1)
	point2 := NewPoint(x2, y2)

	// Вычисляем расстояние между точками
	distance := point1.Distance(point2)

	// Выводим результат
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}

// Point представляет точку в двумерном пространстве с координатами x и y.
type Point struct {
	x, y float64
}

// NewPoint создает новый экземпляр точки.
func NewPoint(x, y float64) *Point {
	return &Point{
		x,
		y,
	}
}

// Distance вычисляет расстояние между текущей точкой и другой точкой.
func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
