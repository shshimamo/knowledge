package utils

import (
	"errors"
	"strconv"
)

func StringSliceToInt64Slice(strings []string) ([]int64, error) {
	ints := make([]int64, len(strings))
	for i, v := range strings {
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, errors.New("failed to convert string slice to int slice")
		}
		ints[i] = num
	}
	return ints, nil
}
