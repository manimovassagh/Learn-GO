package arrays

import (
	"fmt"
)

func PrintArray() {
	// arr := [4]string{"mani", "sahar", "some", "mehdi"}
	// t := arr[3]
	// l := arr[2]
	// a := arr[1]
	// log.Println(arr)
	// log.Println(t)
	// log.Println(l)
	// log.Println(a)
	// log.Println(len(arr))
	sliceSample := []string{"sahar"}
	sliceSample=append(sliceSample, "mani")
	fmt.Println(sliceSample)

}
