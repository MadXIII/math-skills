package main

import (
	"fmt"
	"log"

	"github.com/madxiii/math-skills/calculate"
	"github.com/madxiii/math-skills/get"
)

func main() {
	slice, err := get.Slice()
	if err != nil {
		log.Fatal(err)
	}

	average := calculate.Average(slice)

	median := calculate.Median(slice)

	variance := calculate.Variance(slice)

	deviation := calculate.Deviation(slice)

	fmt.Printf("Average: %d\nMedian: %d\nVariance: %d\nStandard Deviation: %d\n", average, median, variance, deviation) // Print in special format
}
