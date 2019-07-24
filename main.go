package main

import (
	"fmt"
)

func main() {
	fmt.Println("is it go ?")
	foo()

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func foo() {
	fmt.Println("was jetzt ?")
}
