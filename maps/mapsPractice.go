package maps

import "fmt"

type MapStruct struct {
	KeyMap, ValueMap float64
}

var m map[string]MapStruct

func MapLearning() {
	m = make(map[string]MapStruct)
	m["Bell Labs"] = MapStruct{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
