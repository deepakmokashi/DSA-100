package main

import "fmt"

func main() {
	num := 153
	p := num
	temp := 0

	for num != 0 {
		rem := num % 10
		num = num / 10
		temp += rem * rem * rem
	}

	if p == temp {
		fmt.Println("Given number is armstrong number")
	} else {
		fmt.Println("Given number is not armstrong number")
	}
}
