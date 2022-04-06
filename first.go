package main

import (
	"fmt"
	"math"
)

func main() {
	var xlim1 float64 = 2.0
	var xlim2 float64 = 4.0
	var Ntochok float64 = 1000.0
	var f func(float64) = expression
	var result float64 = Simpson(f, xlim1, xlim2, Ntochok)
	fmt.Println(result)
}

func expression(x float64) { return (math.Exp(x)*math.Sin(3.0*math.Pow(x, 2.0)-x) + math.Log(x+2.0)) }

func Simpson(function func(float64), lim1, lim2, n float64) float64 {
	// function - підінтегральна функція
	// lim1 lim2 - межі інтегрування
	// n - кількість кроків
	// обчислимо крок інтегрування

	h := (lim2 - lim1) / n
	// Підсумовуємо всі значення функції множені на відповідні коефіцієнти
	// та множимо суму на крок
	firstPart := function(lim1) + function(lim2)
	//secondPart := function(lim1 + h*i)
	var numbers [n]float64
	for i := 1; i < n; i++ {
		numbers[i] = i
	}

	thirdPart := function(lim1 + h*i)
	var Int float64 = h / 3 * (firstPart + secondPart + thirdPart)

	return Int
}
