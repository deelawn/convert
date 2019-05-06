package convert

import (
	"errors"
	"math"
	"strconv"
)

// UintToBool -- Converts a value from uint64 to boolean
func UintToBool(value uint64) (bool, error) {

	if uint64(1) == value {
		return true, nil
	} else if uint64(0) == value {
		return false, nil
	}

	return false, errors.New(cannotConvertErrMsg)
}

// UintToFloat32 -- Converts a value from uint64 to float32
func UintToFloat32(value uint64) (float32, error) {

	result := float32(value)

	if value > float32MaxPreciseValue {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToFloat64 -- Converts a value from uint64 to float64
func UintToFloat64(value uint64) (float64, error) {

	result := float64(value)

	if value > float64MaxPreciseValue {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToInt -- Converts a value from uint64 to int
func UintToInt(value uint64) (int, error) {

	result := int(value)

	// Need to explicitly check against min/max values because comparing an int64 to uint64
	// with the same bits will have equal values if casting the int to uint64 to do the comparison.
	if uint64(result) != value || (strconv.IntSize == 64 && value > uint64(math.MaxInt64)) ||
		(strconv.IntSize == 32 && value > uint64(math.MaxInt32)) {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToInt8 -- Converts a value from uint64 to int8
func UintToInt8(value uint64) (int8, error) {

	result := int8(value)

	if uint64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToInt16 -- Converts a value from uint64 to int16
func UintToInt16(value uint64) (int16, error) {

	result := int16(value)

	if uint64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToInt32 -- Converts a value from uint64 to int32
func UintToInt32(value uint64) (int32, error) {

	result := int32(value)

	if uint64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToInt64 -- Converts a value from uint64 to int64
func UintToInt64(value uint64) (int64, error) {

	result := int64(value)

	// Need to explicity check against min/max values because comparing an int64 to uint64
	// with the same bits will have equal values if casting the int64 to uint64 to do the comparison.
	if uint64(result) != value || value > uint64(math.MaxInt64) {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToDefaultUint -- Converts a value from uint64 to uint
func UintToDefaultUint(value uint64) (uint, error) {

	result := uint(value)

	if uint64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToUint8 -- Converts a value from uint64 to uint8
func UintToUint8(value uint64) (uint8, error) {

	result := uint8(value)

	if uint64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToUint16 -- Converts a value from uint64 to uint16
func UintToUint16(value uint64) (uint16, error) {

	result := uint16(value)

	if uint64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToUint32 -- Converts a value from uint64 to uint32
func UintToUint32(value uint64) (uint32, error) {

	result := uint32(value)

	if uint64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// UintToString -- Converts a value from uint64 to string
func UintToString(value uint64) (string, error) {

	return strconv.FormatUint(value, 10), nil
}
