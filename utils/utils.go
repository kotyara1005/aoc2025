package utils

import "strconv"

func Atoi(val string) int {
	rv, err := strconv.Atoi(val)
	if err != nil {
		panic(err.Error())
	}
	return rv
}
