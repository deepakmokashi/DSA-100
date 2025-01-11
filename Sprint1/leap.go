package main

import "fmt"

func main() {
	year := 2000

	if year%100 == 0 {
		if year%400 == 0 {
			fmt.Println("Given is leap year ", year)
		} else {
			fmt.Println("Given is not leap year ", year)
		}
	} else if year%4 == 0 {
		fmt.Println("Given is leap year ", year)
	} else {
		fmt.Println("Given is not leap year ", year)
	}
}
