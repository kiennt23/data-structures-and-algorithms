package main

import "fmt"

func main() {
	square, cube := powerSeries(3)
	fmt.Println("Square ", square, "Cube ", cube)
}

func powerSeries(num int) (int, int) {
	return num * num, num * num * num
}
