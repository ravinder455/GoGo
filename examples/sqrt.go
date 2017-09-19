package main

import ("fmt"; "math")

func main() {
	var i = 16.0
	fmt.Println("Inside sqrt function")
	// calcPi()
	// calcSqrt(i)
	a, b := calc(i)
	fmt.Println(a)
	fmt.Println(b)
	
}

func calcPi() {
	fmt.Println("Pi value is ", math.Pi)
}

func calcSqrt(i float64) {
	fmt.Println("Sqrt of 16 is ", math.Sqrt(i))
}

func calc(i float64) (a, b float64) {
	a = math.Pi
	b = math.Sqrt(i)
	return 
}