package loop

import "fmt"

func LoopPractice() {
	for n := 0; n < 10; n++ {
		fmt.Println("Loop Is going ", n, " Times")
	}
}

func AnotherForLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
