package function

import (
	"fmt"
	"math"
)

type dualError struct {
	Num     float64
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("Woring!!!, because \"%f\" is a negative number", e.Num)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f}
	}
	return math.Sqrt(f), nil
}
