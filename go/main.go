package main

import (
	"fmt"
	"math"

	"integral/src/simpson"
)

func main() {
	xlim1 := 2.0
	xlim2 := 4.0
	Ntochok := 1000.0
	//var f func(float64) = expression
	f := func(x float64) float64 {
		return math.Exp(x)*math.Sin(3.0*math.Pow(x, 2.0)-x) + math.Log(x+2.0)
	}
	var result float64 = simpson.Simpson(f, xlim1, xlim2, Ntochok)
	fmt.Println("Result = ", result)
}

/*
func Simpson(function func(x float64) float64, lim1, lim2, n float64) float64 {
	// function - підінтегральна функція
	// lim1 lim2 - межі інтегрування
	// n - кількість кроків
	// обчислимо крок інтегрування

	var h float64 = (lim2 - lim1) / n
	fmt.Println("h=", h)
	// Формулу можна умовно поділити на 3 частини
	// Перша - сума першого та останнього значення підінтегральної функції
	// Друга - сума членів з парними індексами, що множиться на 2
	// Третя - сума членів з непарними індексами, що множиться на 4
	firstPart := function(lim1) + function(lim2)
	// fmt.Println("firstPart = ", firstPart)
	// Заповнюємо масив("зріз") точок
	var numbers []int = make([]int, int(n))
	for i := 1; i < int(n); i++ {
		numbers[i] = i
	}

	// Розподіяємо на парні та непарні
	var unpair, pair []int
	for _, i := range numbers[1 : len(numbers)-1] {
		if i%2 == 0 {
			pair = append(pair, i)
		} else {
			unpair = append(unpair, i)
		}
	}

	// Частина друга ==============================================================================
	// Обраховуємо значення функції в відповідних точках
	var pairIndexValue []float64
	for _, i := range pair {
		pairIndexValue = append(pairIndexValue, function(lim1)+h*float64(i))
	}
	fmt.Println("+++++++++++++=======", function(lim1)+h*float64(pair[0]))
	// Сумуємо
	var pairTotal float64
	for _, i := range pairIndexValue {
		pairTotal += float64(i)
	}
	secondPart := 2 * pairTotal
	// fmt.Println("secondPart", secondPart)

	// Третя частина     =========================================================================
	var unpairIndexValue []float64
	for _, i := range pair {
		unpairIndexValue = append(unpairIndexValue, function(float64(lim1)+float64(h)*float64(i)))
	}
	// Сумуємо
	var unpairTotal float64
	for _, i := range pairIndexValue {
		unpairTotal += float64(i)
	}

	// fmt.Println("pair = ", pairIndexValue)
	// fmt.Println("unpair = ", unpairIndexValue)

	thirdPart := 4 * unpairTotal
	// fmt.Println("thirdPart = ", thirdPart)
	//  Підраховуємо все разом ===================================================================
	var simp float64 = h / 3.0 * (firstPart + secondPart + thirdPart)

	return simp
}
*/
