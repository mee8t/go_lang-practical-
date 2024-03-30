package main

import "fmt"

func main() {
	num := []int{10, 20, 30, 40, 50}

	fmt.Println("the array is", num)

	sum := 0
	for _num := range num {
		sum += _num
	}
	fmt.Printf("the sum of array is:%d \n", sum)
	fmt.Printf("the length of array  %d \n", len(num))
	fmt.Printf("value of index 0: %d\n ", num[0])
	fmt.Printf("value of index 1: %d\n ", num[1])
	fmt.Printf("value of index 2: %d\n ", num[2])
	fmt.Printf("value of index 3: %d\n ", num[3])
	fmt.Printf("value of index 4: %d\n ", num[4])
}
