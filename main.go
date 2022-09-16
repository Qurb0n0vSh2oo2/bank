package main

import (
	"fmt"
	"math"
)

// func main() {
// 	var a, x int
// 	fmt.Scan(&a)
// 	x = fibonacci(a)
// 	fmt.Println(x)
// }
// func fibonacci(n int) int {
// 	if(n <= 2) {
// 		return 1
// 	}
// 	return fibonacci(n - 1) + fibonacci(n - 2);
// }
func main() {
	c := Circle{0, 0, 5}
    fmt.Println(circleArea(c))
}
func circleArea(c Circle) float64 {
    return math.Pi * c.r * c.r
}