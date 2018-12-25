package main

import (
	"fmt"
)

func main() {
	arr2 := []int{2, 3, 1, 0, -99, -2, 77, 74}
	output := insertionsort(arr2)
	fmt.Println(output)
}

//I am using logging methods to understand each step of the algorithm
//I find it easier to visualize in my head how it is operating on the numbers

func insertionsort(arr []int) []int {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i > -1 && arr[i] > key {
			arr[i+1] = arr[i]
			i = i - 1
			fmt.Println("")
			fmt.Println(arr)
		}
		fmt.Println("I am out of the inner loop")
		fmt.Println("")
		fmt.Printf("-- j - 1 (or i) is %v and the key is %v --\n", i, key)
		arr[i+1] = key
		fmt.Println("bottom of loop", arr)
	}
	return arr
}
