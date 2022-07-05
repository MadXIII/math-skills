package get

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	errEmpty error = errors.New("empty file")
	errArgs  error = errors.New("wrong input arguments")
)

func Slice() ([]float64, error) {
	var slice []float64
	arg := os.Args[1:]

	if len(arg) != 1 {
		return slice, errArgs
	}

	file := arg[0]

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return slice, fmt.Errorf("error: %v, while reading file: %s", err, file)
	}

	slice, err = convert(b)
	if err != nil {
		return slice, fmt.Errorf("error while converting: %s", err)
	}

	return slice, nil
}

func convert(b []byte) ([]float64, error) {
	var temp int
	var err error
	var res []float64

	if len(b) == 0 {
		return res, errEmpty
	}

	splited := strings.Split(string(b), "\n")

	for i := range splited {
		if splited[i] == "" {
			continue
		}
		temp, err = strconv.Atoi(splited[i])
		if err != nil {
			return res, err
		}
		res = append(res, float64(temp))
	}

	return res, nil
}
