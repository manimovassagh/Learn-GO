package variables

import (
	"fmt"
)

func needFloat(x float64) float64 {
	return x * 0.1
}
func TypeConvertion() {
	var i int = 42
	var f float64 = float64(i)
	g := 0.867 + 0.5i
	fmt.Println(f)
	fmt.Println(g)
	fmt.Printf("g is of type %T\n", g)
	const rabbat = 13.5
	const big = 1 << 100
	fmt.Printf("rabbat is of type %T\n", rabbat)
	fmt.Printf("rabbat value is %v\n", rabbat)
	fmt.Println(needFloat(big))
}
