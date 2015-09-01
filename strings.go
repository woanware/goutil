package goutil

import (
	"strings"
	"fmt"
)

// Emulates the python partition function
func Partition(data string, separator string) (pre string, post string) {
	index := strings.Index(data, separator)
	if index == -1 {
		return "", ""
	}

	return data[:index], data[index+1:]
}

//
func GetSeparator(s string) rune {
	var sep string
	s = `"` + s + `"`
	fmt.Sscanf(s, "%q", &sep)

	return ([]rune(sep))[0]
}
