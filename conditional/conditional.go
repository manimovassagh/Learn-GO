package conditional

import (
	"fmt"
	"math"
)

func Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	fmt.Println(lim)
	return lim
}
