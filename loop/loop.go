package loop

import "fmt"

func LoopPractice() {
	for n := 0; n < 10; n++ {
		fmt.Println("Loop Is going ", n, " Times")
	}
}

// var v = []string{"mani","sahar"}
func PrintSliceVim() {
var t =[]string{"mani","sahar","neovim!!!"}
	var strSlice = []string{"India", "Canada", "Japan"}
	for _, v := range strSlice {
		fmt.Println(v)
	}

  for _,l:= range t {
   fmt.Println("I Print from Neovim", l) 
  }
}

func AnotherForLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
