package main

import (
	"RelevantDarkUpgrade/average"
	"fmt"
	"log"
)

func print(val float64, err error) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

func main() {

	even := &average.Even{Values: []float64{1, 3, 4.2, 6, 8, 10}}
	odd := &average.Odd{Values: []float64{1, 4, 3,7, 6, 8, 10}}
	evenodd := &average.EvenOdd{Values: []float64{1, 3, 2.1, 4, 6, 8, 10}}

	fmt.Print("even= ")
	print(even.CalcuateAverage())
	fmt.Println("count of even no=",even.Num)
	fmt.Print("odd= ")
	print(odd.CalcuateAverage())
	fmt.Println("count of odd no=",odd.Num)
	fmt.Print("EvenOdd= ")
	print(evenodd.CalcuateAverage())
	fmt.Println("count of numbers=",evenodd.Num)

}