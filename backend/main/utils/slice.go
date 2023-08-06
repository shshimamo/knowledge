package utils

import (
	"errors"
	"strconv"
)

func IntSliceToInt64Slice(ints []int) []int64 {
	int64s := make([]int64, len(ints))
	for i, v := range ints {
		int64s[i] = int64(v)
	}
	return int64s
}

func StringSliceToIntSlice(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, v := range strings {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("failed to convert string slice to int slice")
		}
		ints[i] = num
	}
	return ints, nil
}
