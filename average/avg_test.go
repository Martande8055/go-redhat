package average

import "testing"

func TestEvenCalcuateAverage(t *testing.T) {

	input := []struct {
		name      string
		val       []float64
		expResult float64
		expError  string
	}{
		{"NotContainEvenNum", []float64{1.0, 3.0, 5.0}, 0, "no even numbers"},
		{"OnlyEvenNumbers", []float64{2.0, 4.0, 6.0}, 4.0, "no"},
		{"AllNumbers", []float64{1.0, 2.0, 3.0, 4.0}, 3.0, "no"},
		{"IncludeNegativeNum", []float64{-2.0, -4.0, 1.0}, -3.0, "no"},
	}

	for _, in := range input {
		t.Run(in.name, func(t *testing.T) {
			e := Even{Values: in.val}
			result, err := e.CalcuateAverage()
			if (err != nil) && in.expError != "no even numbers" {
				t.Fatalf("expected error: %v, got: %v", in.expError, err)
			}
			if result != in.expResult && in.expError == "no" {
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
		{"NotContainOddNum", []float64{2.0, 4.0, 6.0}, 0, "no odd numbers"},
		{"OnlyOddNumbers", []float64{1.0, 3.0, 5.0}, 3.0, "no"},
		{"AllNumbers", []float64{1.0, 2.0, 3.0, 4.0}, 2.0, "no"},
		{"IncludeNegativeNum", []float64{-2.0, -4.0, 1.0}, 1.0, "no"},
	}

	for _, in := range input {
		t.Run(in.name, func(t *testing.T) {
			e := Odd{Values: in.val}
			result, err := e.CalcuateAverage()
			if (err != nil) && in.expError != "no odd numbers" {
				t.Fatalf("expected error: %v, got: %v", in.expError, err)
			}
			if result != in.expResult && in.expError == "no" {
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
		{"SingleNumber", []float64{}, 0, "no numbers"},
		{"SingleNumber", []float64{5.0}, 5.0, "no"},
		{"MultipleNumbers", []float64{1.0, 2.0, 3.0, 4.0}, 2.5, "no"},
		{"WithNegativeNumbers", []float64{-2.0, 3.0, 4.0, 6.0}, 2.75, "no"},
	}

	for _, in := range input {
		t.Run(in.name, func(t *testing.T) {
			e := EvenOdd{Values: in.val}
			result, err := e.CalcuateAverage()
			if (err != nil) && in.expError != "no numbers" {
				t.Fatalf("expected error: %v, got: %v", in.expError, err)
			}
			if result != in.expResult && in.expError == "no" {
				t.Errorf("expected: %v, got: %v", in.expResult, result)
			}
		})
	}

}
