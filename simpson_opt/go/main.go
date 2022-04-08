package main

import (
	"fmt"
	"integral/go/simpson"
	"math"
)

func main() {
	// var gw sync.WaitGroup // Объявление переменной WaitGroup
	// channel := make(chan float64)
	var xlim1 float64 = 2.0
	var xlim2 float64 = 4.0
	var Ntochok int = 1000
	//var f func(float64) = expression
	f := func(x float64) float64 {
		return math.Exp(x)*math.Sin(3.0*math.Pow(x, 2.0)-x) + math.Log(x+2.0)
	}
	var result float64 = simpson.SimpsonGor(f, xlim1, xlim2, Ntochok)
	fmt.Println("Result = ", result)
	// fmt.Scanln()
	// gw.Wait()
}
