package average

import (
	"GO_PROJECT/constants"
	"errors"
	"fmt"
)

type Avg interface {
	CalcuateAverage() (float64, error)
}

type Even struct {
	Num    int
	Values []float64
}

type Odd struct {
	Num    int
	Values []float64
}

type EvenOdd struct {
	Num    int
	Values []float64
}

func (e *Even) CalcuateAverage() (float64, error) {
	var sum float64

	for _, value := range e.Values {
		if value != float64(int64(value)) {
			continue
		}
		if int(value)&1 == 0 {
			sum += value
			e.Num++
		}
	}
	if e.Num == 0 {
		return 0, errors.New(constants.NoEven)
	}
	return sum / float64(e.Num), nil
}

func (o *Odd) CalcuateAverage() (float64, error) {
	var sum float64
	for _, value := range o.Values {
		if value != float64(int64(value)) {
			continue
		}
		if int(value)&1 != 0 {
			sum += value
			o.Num++
		}
	}
	if o.Num == 0 {
		return 0, fmt.Errorf(constants.NoOdd)
	}
	return sum / float64(o.Num), nil
}

func (e *EvenOdd) CalcuateAverage() (float64, error) {
	var sum float64
	for _, value := range e.Values {
		sum += value
		e.Num++
	}
	if e.Num == 0 {
		return 0, errors.New(constants.NoNum)
	}
	return sum / float64(e.Num), nil
}
