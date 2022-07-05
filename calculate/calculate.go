package calculate

import (
	"math"
	"sort"
)

func Average(slice []float64) int {
	var res int
	var temp float64
	var length float64 = float64(len(slice))

	for i := range slice {
		temp += slice[i]
	}
	temp = temp / length

	res = int(math.Round(temp))

	return res
}

func Median(slice []float64) int {
	var temp float64
	var i int = len(slice) / 2 // get the pivot of the slice of numbers

	sort.Float64s(slice)

	if len(slice)%2 == 0 { // is number of slice is even?
		temp = (slice[i-1] + slice[i]) / 2
		return int(math.Round(temp))
	}

	return int(slice[i])
}

func Variance(slice []float64) int {
	var res float64

	ss := sumOfSquares(slice)

	res = ss / float64(len(slice))

	res = math.Round(res)

	return int(res)
}

func Deviation(slice []float64) int {
	var res float64

	ss := sumOfSquares(slice)

	res = ss / float64(len(slice))

	res = math.Sqrt(res)

	res = math.Round(res)

	return int(res)
}

// sumOfSquares - get sum of squares of elements in O(n)^2
func sumOfSquares(slice []float64) float64 {
	var mean float64 // sum of slice / elements
	var temp float64
	var res float64
	var length float64 = float64(len(slice))

	for i := range slice {
		mean += slice[i]
	}

	mean /= length

	for i := range slice {
		temp = slice[i] - mean
		temp *= temp
		res += temp
	}

	return res
}
