package defering

import "fmt"

func DeferSomeValues() {
	// Always started from last defer
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	defer fmt.Println("defer 4")

	fmt.Println("world")

}
