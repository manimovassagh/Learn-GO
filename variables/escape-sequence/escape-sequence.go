package escapesequence

import "fmt"

var (
	customer = "Mani Movassagh"
	c        = "sahar"
)

var (
	t = "mani"
)

func EscapeSequence() {

	fmt.Println("Hi Mani\nthis is new Line\nand this is another line")
	fmt.Printf("Hi: %v", customer)
	fmt.Printf("\nHi: %v", c)
	fmt.Printf("\nHi: %v", t)
}
