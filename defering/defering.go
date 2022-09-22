package defering

import "fmt"

func DeferSomeValues() {

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("world")

}
