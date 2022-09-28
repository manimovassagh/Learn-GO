package rangepack

import "log"

func RangeLearning() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, m := range pow {
		log.Println(i, m)
	}
}
