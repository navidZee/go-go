package main

import "fmt"

func Compare(num1 int, num2 int) (isEqual bool, difference int) {
	if num1 > num2 {
		isEqual = false
		difference = num1 - num2
	} else if num2 > num1 {
		isEqual = false
		difference = num2 - num1
	} else {
		isEqual = true
		difference = 0
	}
	return
}

func main() {
	fmt.Println(Compare(1, 1))
	fmt.Println(Compare(1, 2))
	fmt.Println(Compare(1, 0))
}
