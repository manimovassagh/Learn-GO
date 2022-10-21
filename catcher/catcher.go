package catcher

import "log"

func Error(err error) {
	log.Printf("Some Error is:\n")
	log.Println(err)
}
