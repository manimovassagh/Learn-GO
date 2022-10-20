package arrays

import "log"

func PrintArray() {
	arr := [4]string{"mani", "sahar","some","mehdi"}
	t := arr[3]
	l := arr[2]
	a := arr[1]
	log.Println(arr)
	log.Println(t)
	log.Println(l)
	log.Println(a)
	log.Println(len(arr))
}
