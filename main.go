package main

import "fmt"

func main() {
	// explicit declaration of arrays
	var integerArray1 [3]int
	integerArray1[0] = 1
	integerArray1[1] = 2
	integerArray1[2] = 3
	fmt.Println(integerArray1)

	// implicit declaration of arrays
	integerArray2 := [3]int{1, 2, 3}
	fmt.Println(integerArray2)
}
