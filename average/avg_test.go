package average

import (
	"GO_PROJECT/constants"
	"testing"
)

func TestEvenCalcuateAverage(t *testing.T) {

	input := []struct {
		name      string
		val       []float64
		operation string
		expResult float64
		expError  string
	}{
		{"NotContainEvenNum", []float64{1.0, 3.0, 5.0}, "sum", 0, constants.NoEven},
		{"OnlyEvenNumbers", []float64{2.0, 4.0, 6.0}, "avg", 4.0, constants.NoError},
		{"AllNumbers", []float64{1.0, 2.0, 3.0, 4.0}, "count", 3.0, constants.NoError},
		{"IncludeNegativeNum", []float64{-2.0, -4.0, 1.0}, "avg", -3.0, constants.NoError},
	}

	for _, in := range input {
		t.Run(in.name, func(t *testing.T) {
			// e := Even{Values: in.val}
			var e Avg = &Even{Op: in.operation, Values: in.val}
			result, err := e.Calcuate()
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
		operation string
		expResult float64
		expError  string
	}{
		{"NotContainOddNum", []float64{2.0, 4.0, 6.0}, "sum", 0, constants.NoOdd},
		{"OnlyOddNumbers", []float64{1.0, 3.0, 5.0}, "avg", 3.0, constants.NoError},
		{"AllNumbers", []float64{1.0, 2.0, 3.0, 4.0}, "count", 2.0, constants.NoError},
		{"IncludeNegativeNum", []float64{-2.0, -4.0, 1.0}, "avg", 1.0, constants.NoError},
	}

	for _, in := range input {
		t.Run(in.name, func(t *testing.T) {
			// e := Odd{Values: in.val}
			var e Avg = &Odd{Op: in.operation, Values: in.val}
			result, err := e.Calcuate()
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
		operation string
		expResult float64
		expError  string
	}{
		{"SingleNumber", []float64{}, "sum", 0, constants.NoNum},
		{"SingleNumber", []float64{5.0}, "count", 5.0, constants.NoError},
		{"MultipleNumbers", []float64{1.0, 2.0, 3.0, 4.0}, "sum", 2.5, constants.NoError},
		{"WithNegativeNumbers", []float64{-2.0, 3.0, 4.0, 6.0}, "avg", 2.75, constants.NoError},
	}

	for _, in := range input {
		t.Run(in.name, func(t *testing.T) {
			//e :=  EvenOdd{Values: in.val}
			var e Avg = &EvenOdd{Op: in.operation, Values: in.val}
			result, err := e.Calcuate()
			if (err != nil) && in.expError != constants.NoNum {
				t.Fatalf("expected error: %v, got: %v", in.expError, err)
			}
			if result != in.expResult && in.expError == constants.NoError {
				t.Errorf("expected: %v, got: %v", in.expResult, result)
			}
		})
	}

}

func BenchmarkEvenCalcuateAverage(b *testing.B) {

	var even Avg = &Even{Values: []float64{1.0, 2.0, 3.0, 4.0}}
	for i := 0; i < b.N; i++ {
		even.Calcuate()
	}
}

func BenchmarkOddCalcuateAverage(b *testing.B) {
	var odd Avg = &Odd{Values: []float64{1.0, 2.0, 3.0, 4.0}}

	for i := 0; i < b.N; i++ {
		odd.Calcuate()
	}
}

func BenchmarkEvenOddCalcuateAverage(b *testing.B) {
	var evenodd Avg = &EvenOdd{Values: []float64{1.0, 2.0, 3.0, 4.0}}
	for i := 0; i < b.N; i++ {
		evenodd.Calcuate()
	}
}
