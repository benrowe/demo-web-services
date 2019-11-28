package transformations

import "strconv"

const (
	intConversionBase = 10
)

func uint32ToStr(v uint) string {
	return strconv.FormatUint(uint64(v), intConversionBase)
}
