package simpson

func Simpson(function func(x float64) float64, lim1, lim2 float64, n int) float64 {
	// function - підінтегральна функція
	// lim1 lim2 - межі інтегрування
	// n - кількість кроків
	// обчислимо крок інтегрування
	var h float64 = (lim2 - lim1) / float64(n)

	// Формулу можна умовно поділити на 3 частини
	// Перша - сума першого та останнього значення підінтегральної функції
	// Друга - сума членів з парними індексами, що множиться на 2
	// Третя - сума членів з непарними індексами, що множиться на 4
	firstPart := function(lim1) + function(lim2)

	// Розподіяємо на парні та непарні й одразу обчислюємо
	var pairIndexValue, oddIndexValue []float64
	for i := 1; i < int(n); i++ {
		if i%2 == 0 {
			pairIndexValue = append(pairIndexValue, function(lim1+h*float64(i)))
		} else {
			oddIndexValue = append(oddIndexValue, function(lim1+h*float64(i)))
		}
	}

	// Сумуємо
	var pairTotal, unpairTotal float64
	for _, i := range pairIndexValue {
		pairTotal += float64(i)
	}
	for _, i := range oddIndexValue {
		unpairTotal += float64(i)
	}

	//  Підраховуємо все разом ===================================================================
	var simp float64 = h / 3.0 * (firstPart + pairTotal*2 + unpairTotal*4)

	return simp
}
