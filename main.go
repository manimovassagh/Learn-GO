package main

import (
	constantlearn "app/constant-learn"
	"log"

	sgeht "github.com/manimovassagh/Sgeht"
)

func main() {
	// variables.Variables()
	// variables.DefaultVals()
	// variables.TypeConvertion()
	// loop.LoopPractice()
	// loop.AnotherForLoop()
	// conditional.Pow(4, 3, 12)
	// switcher.SwitchingData()
	// switcher.SwitchDaysInWeek()
	//pointers.PointerPractice()
	// switcher.SwitchingData()
	// defering.DeferSomeValues()
	// defering.StackDeferItems()
	//pointerspack.LearnPointersOne()
	//pointers.PointerPractice()
	//rangepack.RangeLearning()
	//maps.MapLearning()
	//m := functions.FuncLearnA(functions.Multipling)
	constantlearn.CheckIota()
	r := sgeht.SReq("https://jsonplaceholder.typicode.com/posts")
	log.Println(r)

}
