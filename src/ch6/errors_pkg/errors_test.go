package errors_pkg

import (
	"errors"
	"math"
	"testing"
)

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
func TestSqrt(t *testing.T) {
	result, err := Sqrt(-10)
	t.Log(result, err)
	result, err = Sqrt(10)
	t.Log(result, err)
}
