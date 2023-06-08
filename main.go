package main

import "fmt"

func main() {
	fmt.Println(CalcSum(1800, 1.5), "so`m")
	fmt.Println(Fibo(5))
}

func CalcSum(price float64, kg float64) float64 {
	return price * kg
}

func Fibo(num int) int {
	return num
}
