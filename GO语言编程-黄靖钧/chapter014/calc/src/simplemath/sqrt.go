package simplemath

import (
	"fmt"
	"math"
)

func Sqrt(i int) string {
	v := math.Sqrt(float64(i))
	return fmt.Sprintf("%.4f", v)
}
