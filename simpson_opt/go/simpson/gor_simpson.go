package simpson

import (
	"sync"
)

func SimpsonGor(function func(x float64) float64, lim1, lim2 float64, n int) float64 {
	var wg sync.WaitGroup
	// function - підінтегральна функція
	// lim1 lim2 - межі інтегрування
	// n - кількість кроків
	// обчислимо крок інтегрування

	var h float64 = (lim2 - lim1) / float64(n)
	// Формулу Сімпсона можна поділити на 3 частини
	// Перша - сума першого та останнього значення підінтегральної функції
	// Друга - сума членів з парними індексами, що множиться на 2
	// Третя - сума членів з непарними індексами, що множиться на 4

	// Заповнюємо масив("зріз") точок
	// var numbers []int = make([]int, int(n))
	var pairIndexValue, oddIndexValue []float64
	for i := 1; i < int(n); i++ {
		if i%2 == 0 {
			pairIndexValue = append(pairIndexValue, function(lim1+h*float64(i)))
		} else {
			oddIndexValue = append(oddIndexValue, function(lim1+h*float64(i)))
		}
	}
	// Сумуємо на всіх етапах
	wg.Add(3)
	firstPartCh := make(chan float64)
	pairChTot := make(chan float64)
	oddChTot := make(chan float64)

	go FirstPart(function, lim1, lim2, firstPartCh, &wg)
	go SumTotal(pairIndexValue, pairChTot, &wg)
	go SumTotal(oddIndexValue, oddChTot, &wg)

	firstPart := <-firstPartCh
	pairTotal := <-pairChTot
	oddTotal := <-oddChTot

	// Збираємо все разом
	var simp float64 = h / 3.0 * (firstPart + pairTotal*2 + oddTotal*4)
	wg.Wait()

	return simp
}

func PairOdd1(arr []int, pairCh, oddCh chan int) {
	// var odd, pair []int
	defer close(pairCh)
	defer close(oddCh)
	for _, i := range arr[1:] {
		if i%2 == 0 {
			pairCh <- i
		} else {
			oddCh <- i
		}
	}
}
func SumTotal(arr []float64, ch chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	oddTotal := 0.0
	for _, i := range arr {
		oddTotal += float64(i)
	}
	ch <- oddTotal
}
func FirstPart(function func(x float64) float64, lim1, lim2 float64, ch chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	firstPart := function(lim1) + function(lim2)
	ch <- firstPart
}
