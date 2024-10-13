package average

import (
	"GO_PROJECT/constants"
	"errors"
	"fmt"
)

type Avg interface {
	Calcuate() (float64, error)
}

type Even struct {
	Num    int
	Op     string
	Values []float64
}

type Odd struct {
	Num    int
	Op     string
	Values []float64
}

type EvenOdd struct {
	Num    int
	Op     string
	Values []float64
}

func (e *Even) Calcuate() (float64, error) {
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
	switch e.Op {
	case "avg":
		return sum / float64(e.Num), nil
	case "sum":
		return sum, nil
	case "count":
		return float64(e.Num), nil
	}
	return 0, errors.New("invalid Operation")
}

func (o *Odd) Calcuate() (float64, error) {
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
	switch o.Op {
	case "avg":
		return sum / float64(o.Num), nil
	case "sum":
		return sum, nil
	case "count":
		return float64(o.Num), nil
	}
	return 0, errors.New("invalid operation")
}

func (e *EvenOdd) Calcuate() (float64, error) {
	var sum float64
	for _, value := range e.Values {
		sum += value
		e.Num++
	}
	if e.Num == 0 {
		return 0, errors.New(constants.NoNum)
	}
	switch e.Op {
	case "avg":
		return sum / float64(e.Num), nil
	case "sum":
		return sum, nil
	case "count":
		return float64(e.Num), nil
	}
	return 0, errors.New("invalid operation")
}
