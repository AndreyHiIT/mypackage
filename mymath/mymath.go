package mymath

import (
	"math"
)

// Abs возвращает абсолютное значение числа x.
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// Max возвращает максимальное значение среди двух чисел x и y.
func Max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

// Sqrt возвращает квадратный корень числа x.
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Yn вычисляет значение функции Бесселя второго рода.
func Yn(n int, x float64) float64 {
	return math.Yn(n, x) // Используем встроенную функцию для вычисления Jn
}

func Acos(x float64) float64 {
	return math.Acos(x)
}
func Acosh(x float64) float64 {
	return math.Acosh(x)
}
func Asin(x float64) float64 {
	return math.Asin(x)
}

//v1.0.0
