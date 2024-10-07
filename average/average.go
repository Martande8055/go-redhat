package average

import (
	"fmt"
	"os"
)

// CalculateAverage reads float64 numbers from standard input,
// calculates the average, and returns it.
func CalculateAverage() (float64, error) {
	var n int
	var sum float64

	for {
    	var val float64
    	_, err := fmt.Fscan(os.Stdin, &val)
    	if err != nil {
        	break
    	}
    	sum += val
    	n++
	}

	if n == 0 {
    	return 0, fmt.Errorf("no numbers provided")
	}
	return sum / float64(n), nil
}
