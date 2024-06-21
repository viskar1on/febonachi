package logic

import (
	"math"
)

// Проверяет число на квадрат
func IsPerfectSquare(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	return sqrt*sqrt == num
}

// Проверяет является ли число Фибоначчи
func IsFibonacci(num int) bool {
	return IsPerfectSquare(5*num*num+4) || IsPerfectSquare(5*num*num-4)
}

// Возвращает ближайшее число Фибоначчи
func GetNearestFibonacci(num int) int {
	for i := 1; ; i++ {
		if IsFibonacci(num + i) {
			return num + i
		}
		if IsFibonacci(num - i) {
			return num - i
		}
	}
}

// Возвращает соседние числа Фибоначчи
func GetFibonacciNeighbours(num int) (int, int) {
	if num == 0 {
		return 0, 1
	}
	a, b := 0, 1
	for b < num {
		a, b = b, a+b
	}
	return a, b + a
}
