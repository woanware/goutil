package goutil

import (
	"strconv"
)

// ##### Methods #############################################################

// Converts an Int64 to a string
func ConvertInt64ToString(data int64) string {
	return strconv.FormatInt(data, 10 )
}

// Converts an Int64 to a string
func ConvertIntToString(data int) string {
	return strconv.FormatInt(int64(data), 10 )
}

// Converts an Int64 to a string
func ConvertUInt16ToString(data uint16) string {
	return strconv.FormatInt(int64(data), 10 )
}
