package main

import "fmt"

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func sum2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	x := sum(s[:len(s)/2])
	y := sum(s[len(s)/2:])

	z := x + y

	fmt.Println(x, y, z)

	c := make(chan int)
	d := make(chan int)

	go sum2(s[:len(s)/2], c)
	go sum2(s[len(s)/2:], d)

	a, b := <-c, <-d

	fmt.Println(a, b, a+b)

}
