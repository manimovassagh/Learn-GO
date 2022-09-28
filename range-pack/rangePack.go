package rangepack

import "log"

func RangeLearning() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for index, element := range pow {
		log.Println(index, element)
	}
}
