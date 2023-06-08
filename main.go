package main

import "fmt"

func main() {
	fmt.Println(CalcSum(1800, 1.5), "so`m")
}

func CalcSum(price float64, kg float64) float64 {
	return price * kg
}
