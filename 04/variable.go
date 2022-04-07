package main

import (
	"fmt"
)

func Compare(num1 int, num2 int) (isEqual bool, difference int) {

	switch /*x := 2;*/ {
	case num1 > num2:
		/* println(x) */
		isEqual = false
		difference = num1 - num2
	case num2 > num1:
		isEqual = false
		//fallthrough
		difference = num2 - num1
	default:
		isEqual = true
		difference = 0
	}

	return
}

func Factorial(number int) (sum int) {
	sum = 1
	for number > 0 {
		sum = sum * number
		number--
	}
	return
}

func main() {

	//var int_1 int8 = 4   (0 to 255)
	//var int_2 int16 = 4  (0 to 65535)
	//var int_3 int32 = 4 (0 to 4294967295)
	//var int_4 int64 = 4 (0 to 184467440737095516145)

	//var float_1 float32 = 0.1
	//var float_2 float64 = 0.2

	fmt.Println(Compare(1, 1))
	fmt.Println(Compare(2, 1))
	fmt.Println(Compare(1, 1))

	fmt.Println(Factorial(5))
}
