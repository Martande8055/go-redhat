package main

import (
	"GO_PROJECT/average"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func print(val float64, err error) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

func bitArrayTofloat(b []byte) ([]float64, error) {
	data := string(b)

	string_Numbers := strings.Split(data, ",")

	var nums []float64

	for _, str := range string_Numbers {
		str = strings.TrimSpace(str)
		if num, err := strconv.ParseFloat(str, 64); err != nil {

			return []float64{}, err
		} else {
			nums = append(nums, num)
		}
	}
	return nums, nil
}

func main() {

	var odd average.Avg = &average.Odd{Values: []float64{1, 4, 3, 7, 6.1, 8, 10}}
	var evenodd average.Avg = &average.EvenOdd{Values: []float64{1, 3, 2.1, 4, 6, 8, 10}}

	fmt.Print("odd= ")
	print(odd.CalcuateAverage())
	fmt.Println("count of odd no=", odd.(*average.Odd).Num)
	fmt.Print("EvenOdd= ")
	print(evenodd.CalcuateAverage())
	fmt.Println("count of numbers=", evenodd.(*average.EvenOdd).Num)

	http.HandleFunc("/even", evenAvgHandler)
	http.HandleFunc("/odd", oddAvgHandler)
	http.HandleFunc("/evenodd", evenoddAvgHandler)
	log.Println("server started")
	log.Fatal(http.ListenAndServe(":3100", nil))

}

func evenAvgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read req body", http.StatusInternalServerError)
		return
	}

	nums, err := bitArrayTofloat(body)
	if err != nil {
		http.Error(w, "invalid input", http.StatusNotAcceptable)
		return
	}

	var even average.Avg = &average.Even{Values: nums}

	avg, err := even.CalcuateAverage()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Avg of even no is %.2f \n", avg)

}

func oddAvgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read req body", http.StatusInternalServerError)
		return
	}

	nums, err := bitArrayTofloat(body)
	if err != nil {
		http.Error(w, "invalid input", http.StatusNotAcceptable)
		return
	}

	var odd average.Avg = &average.Odd{Values: nums}

	avg, err := odd.CalcuateAverage()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Avg of even no is %.2f \n", avg)

}

func evenoddAvgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read req body", http.StatusInternalServerError)
		return
	}

	nums, err := bitArrayTofloat(body)
	if err != nil {
		http.Error(w, "invalid input", http.StatusNotAcceptable)
		return
	}

	var evenodd average.Avg = &average.EvenOdd{Values: nums}

	avg, err := evenodd.CalcuateAverage()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Avg of even no is %.2f \n", avg)

}
