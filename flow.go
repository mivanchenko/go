package main

import (
	"fmt"
	"math"
)

func main() {
//	sum := 0
//	for i := 0; i < 10; i++ {
//		sum += i
//	}
//	fmt.Println(sum)

//	var sum uint64 = 1
	sum := 1
//	for ; sum < 8; {
	for sum < 9 {
		sum += sum
	}
//	fmt.Println(sum)

	for {
		sum += sum
		break
	}

//	fmt.Println(sqrt(2), sqrt(-4))
//	fmt.Println(
//		pow(3, 2, 10),
//		pow(3, 3, 20),
//	)
	var a, p float64 = 2, 0.01
	fmt.Println("Square root of", a, "with precision", p, "is:", Sqrt(a, p))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
		fmt.Printf("%v >= %v\n", v, lim)
	}
	return lim
}

func Sqrt(x, lim float64) float64 {
	var z float64 = x / 2
	for i := 0; i < 1000000; i++ {
//	for i := 0; i < 1000; i++ {
		D := (z*z - x)
		z -= D / 2*z
		if math.Abs(D) < lim {
			fmt.Println("Delta =", D)
			break
		}
	}
	return z
}
