package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Введите коэффициент a:")
	fmt.Scanln(&a)

	fmt.Println("Введите коэффициент b:")
	fmt.Scanln(&b)

	fmt.Println("Введите коэффициент c:")
	fmt.Scanln(&c)

	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("x1 = %.2f\n", x1)
		fmt.Printf("x2 = %.2f\n", x2)
	} else if discriminant == 0 {
		x := -b / (2 * a)
		fmt.Printf("x = %.2f\n", x)
	} else {
		fmt.Println("Нет действительных корней")
	}
}
