package main

import "fmt"

func main() {
	fmt.Println("counting")
	fmt.Println("\nLast In ---> First Out\n")

	for x := 0; x < 5; x++ {
		defer fmt.Println(x, " popped")
		fmt.Println(x, "deferred")
	}

	fmt.Println("finished\n")
}