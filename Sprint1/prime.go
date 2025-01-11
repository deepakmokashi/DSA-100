// package main

// import "fmt"

// func checkPrime(num int) bool {
// 	temp := true
// 	for i := 2; i < num; i++ {
// 		if num%i == 0 {
// 			temp = false
// 			return temp
// 		}
// 	}
// 	return temp
// }
// func main() {
// 	num := 15
// 	isPrime := checkPrime(num)

// 	if isPrime {
// 		fmt.Printf("num is %d is Prime number", num)
// 	} else {
// 		fmt.Printf("num is %d is not Prime number", num)
// 	}
// }