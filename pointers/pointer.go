package pointers

import (
	"log"
)

func PointerPractice() {
	log.Println("This is pointer")
	p := 12
	i := &p
	*i = 44

	log.Println("p is now ", p)
	log.Println("i is now ", i)
	log.Println("*p is now ", p)
	log.Println("*i ", *i)

}
