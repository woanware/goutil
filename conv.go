package goutil

import (

)

// ##### Methods #############################################################

// Converts an Int64 to a string
func convertInt64ToString(data int64) string {
	return strconv.FormatInt(data, 10 )
}

// Converts an Int64 to a string
func convertIntToString(data int) string {
	return strconv.FormatInt(int64(data), 10 )
}

// Converts an Int64 to a string
func convertUInt16ToString(data uint16) string {
	return strconv.FormatInt(int64(data), 10 )
}
