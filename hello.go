package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var c, python, java bool
var i, j int = 1, 2

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("Hello Міша")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

//	var i int
	var i, j int = 3, 4
//	var c, python, java = true, false, "no!"
	k := 5
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

//func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
