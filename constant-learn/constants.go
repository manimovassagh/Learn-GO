package constantlearn

import (
	"log"
)

const (
	monday    = iota + 1
	tuesday   = iota + 1
	wednesday = iota + 1
	tursday   = iota + 1
	friday    = iota + 1
	saturday  = iota + 1
	sunday    = iota + 1
)

const (
	EST = -(5 + iota)
	_
	MST
	PST
)

func CheckIota() {
	log.Println(monday)
	log.Println(tuesday)
	log.Println(wednesday)
	log.Println(tuesday)
	log.Println(friday)
	log.Println(saturday)
	log.Println(sunday)
	log.Println(EST, MST, PST)

}
