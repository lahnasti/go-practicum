package main

import "fmt"

func main () {
	nums := [5]int{1,2,3,4,5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}