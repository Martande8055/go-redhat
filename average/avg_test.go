package average

import (
	"GO_PROJECT/constants"
	"testing"
)

func TestEvenCalcuateAverage(t *testing.T) {

	input := []struct {
		name      string
		val       []float64
		expResult float64
		expError  string
	}{
		{"NotContainEvenNum", []float64{1.0, 3.0, 5.0}, 0, constants.NoEven},
		{"OnlyEvenNumbers", []float64{2.0, 4.0, 6.0}, 4.0, constants.NoError},
		{"AllNumbers", []float64{1.0, 2.0, 3.0, 4.0}, 3.0, constants.NoError},
		{"IncludeNegativeNum", []float64{-2.0, -4.0, 1.0}, -3.0, constants.NoError},
	}

	for _, in := range input {
		t.Run(in.name, func(t *testing.T) {
			e := Even{Values: in.val}
			result, err := e.CalcuateAverage()
			if (err != nil) && in.expError != constants.NoEven {
				t.Fatalf("expected error: %v, got: %v", in.expError, err)
			}
			if result != in.expResult && in.expError == constants.NoError {
				t.Errorf("expected: %v, got: %v", in.expResult, result)
			}
		})
	}

}

func TestOddCalcuateAverage(t *testing.T) {

	input := []struct {
		name      string
		val       []float64
		expResult float64
		expError  string
	}{
		{"NotContainOddNum", []float64{2.0, 4.0, 6.0}, 0, constants.NoOdd},
		{"OnlyOddNumbers", []float64{1.0, 3.0, 5.0}, 3.0, constants.NoError},
		{"AllNumbers", []float64{1.0, 2.0, 3.0, 4.0}, 2.0, constants.NoError},
		{"IncludeNegativeNum", []float64{-2.0, -4.0, 1.0}, 1.0, constants.NoError},
	}

	for _, in := range input {
		t.Run(in.name, func(t *testing.T) {
			e := Odd{Values: in.val}
			result, err := e.CalcuateAverage()
			if (err != nil) && in.expError != constants.NoOdd {
				t.Fatalf("expected error: %v, got: %v", in.expError, err)
			}
			if result != in.expResult && in.expError == constants.NoError {
				t.Errorf("expected: %v, got: %v", in.expResult, result)
			}
		})
	}

}

func TestAllCalcuateAverage(t *testing.T) {

	input := []struct {
		name      string
		val       []float64
		expResult float64
		expError  string
	}{
		{"SingleNumber", []float64{}, 0, constants.NoNum},
		{"SingleNumber", []float64{5.0}, 5.0, constants.NoError},
		{"MultipleNumbers", []float64{1.0, 2.0, 3.0, 4.0}, 2.5, constants.NoError},
		{"WithNegativeNumbers", []float64{-2.0, 3.0, 4.0, 6.0}, 2.75, constants.NoError},
	}

	for _, in := range input {
		t.Run(in.name, func(t *testing.T) {
			e := EvenOdd{Values: in.val}
			result, err := e.CalcuateAverage()
			if (err != nil) && in.expError != constants.NoNum {
				t.Fatalf("expected error: %v, got: %v", in.expError, err)
			}
			if result != in.expResult && in.expError == constants.NoError {
				t.Errorf("expected: %v, got: %v", in.expResult, result)
			}
		})
	}

}
