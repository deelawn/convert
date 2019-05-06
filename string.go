package convert

import (
	"strconv"
	"strings"
)

// StringToBool -- Converts a value from string to boolean
func StringToBool(value string) (bool, error) {

	return strconv.ParseBool(strings.TrimSpace(value))
}

// StringToFloat32 -- Converts a value from string to float32
func StringToFloat32(value string) (float32, error) {

	result, err := strconv.ParseFloat(strings.TrimSpace(value), 32)

	return float32(result), err
}

// StringToFloat64 -- Converts a value from string to float64
func StringToFloat64(value string) (float64, error) {

	return strconv.ParseFloat(strings.TrimSpace(value), 64)
}

// StringToInt -- Converts a value from string to int
func StringToInt(value string) (int, error) {

	result64, err := strconv.ParseInt(strings.TrimSpace(value), 10, 0)

	return int(result64), err
}

// StringToInt8 -- Converts a value from string to int8
func StringToInt8(value string) (int8, error) {

	result64, err := strconv.ParseInt(strings.TrimSpace(value), 10, 8)

	return int8(result64), err
}

// StringToInt16 -- Converts a value from string to int16
func StringToInt16(value string) (int16, error) {

	result64, err := strconv.ParseInt(strings.TrimSpace(value), 10, 16)

	return int16(result64), err
}

// StringToInt32 -- Converts a value from string to int32
func StringToInt32(value string) (int32, error) {

	result64, err := strconv.ParseInt(strings.TrimSpace(value), 10, 32)

	return int32(result64), err
}

// StringToInt64 -- Converts a value from string to int64
func StringToInt64(value string) (int64, error) {

	return strconv.ParseInt(strings.TrimSpace(value), 10, 64)
}

// StringToUint -- Converts a value from string to uint
func StringToUint(value string) (uint, error) {

	result64, err := strconv.ParseUint(strings.TrimSpace(value), 10, 0)

	return uint(result64), err
}

// StringToUint8 -- Converts a value from string to uint8
func StringToUint8(value string) (uint8, error) {

	result64, err := strconv.ParseUint(strings.TrimSpace(value), 10, 8)

	return uint8(result64), err
}

// StringToUint16 -- Converts a value from string to uint16
func StringToUint16(value string) (uint16, error) {

	result64, err := strconv.ParseUint(strings.TrimSpace(value), 10, 16)

	return uint16(result64), err
}

// StringToUint32 -- Converts a value from string to uint32
func StringToUint32(value string) (uint32, error) {

	result64, err := strconv.ParseUint(strings.TrimSpace(value), 10, 32)

	return uint32(result64), err
}

// StringToUint64 -- Converts a value from string to uint64
func StringToUint64(value string) (uint64, error) {

	return strconv.ParseUint(strings.TrimSpace(value), 10, 64)
}
