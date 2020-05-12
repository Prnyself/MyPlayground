package strings

import "strconv"

func ConvToFloat(s string, bitSize int) (float64, error) {
	return strconv.ParseFloat(s, bitSize)
}
