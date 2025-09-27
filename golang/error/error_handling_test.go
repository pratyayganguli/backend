package error

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func TestCustomError(t *testing.T) {
	singleDigitNumber := 10
	if singleDigitNumber >= 10 {
		e := errors.New("not a single digit number!")
		log.Panic(e)
	}
	log.Print("-------test passed-------")
}

type DivideError struct {
	Dividend float64
	Divisor  float64
}

func (de *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %.2f by %.2f", de.Dividend, de.Divisor)
}

func TestDivisorError(t *testing.T) {
	a := 10.0
	b := 0.0
	if b == 0.0 {
		e := DivideError{
			Dividend: a,
			Divisor:  b,
		}
		log.Panic(e.Error())
	} else {
		log.Printf("%.2f", a/b)
	}
}
