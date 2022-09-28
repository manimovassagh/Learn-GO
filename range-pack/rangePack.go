package rangepack

import "log"

func RangeLearning() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for index, element := range pow {
		log.Println(index, element)
	}
	for index, _ := range pow {
		log.Println("index ",index)
	}
	for _, element := range pow {
		log.Println("element ",element)
	}

}
