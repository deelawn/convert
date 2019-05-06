package convert

import (
	"errors"
	"strconv"
)

// IntToBool -- Converts a value from int64 to boolean
func IntToBool(value int64) (bool, error) {

	if int64(1) == value {
		return true, nil
	} else if int64(0) == value {
		return false, nil
	}

	return false, errors.New(cannotConvertErrMsg)
}

// IntToFloat32 -- Converts a value from int64 to float32
func IntToFloat32(value int64) (float32, error) {

	result := float32(value)

	if value < float32MinPreciseValue || value > float32MaxPreciseValue {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToFloat64 -- Converts a value from int64 to float64
func IntToFloat64(value int64) (float64, error) {

	result := float64(value)

	if value < float64MinPreciseValue || value > float64MaxPreciseValue {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToDefaultInt -- Converts a value from int64 to int
func IntToDefaultInt(value int64) (int, error) {

	result := int(value)

	if int64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToInt8 -- Converts a value from int64 to int8
func IntToInt8(value int64) (int8, error) {

	result := int8(value)

	if int64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToInt16 -- Converts a value from int64 to int16
func IntToInt16(value int64) (int16, error) {

	result := int16(value)

	if int64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToInt32 -- Converts a value from int64 to int32
func IntToInt32(value int64) (int32, error) {

	result := int32(value)

	if int64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToUint -- Converts a value from int64 to uint
func IntToUint(value int64) (uint, error) {

	result := uint(value)

	if int64(result) != value || value < 0 {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToUint8 -- Converts a value from int64 to uint8
func IntToUint8(value int64) (uint8, error) {

	result := uint8(value)

	if int64(result) != value || value < 0 {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToUint16 -- Converts a value from int64 to uint16
func IntToUint16(value int64) (uint16, error) {

	result := uint16(value)

	if int64(result) != value || value < 0 {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToUint32 -- Converts a value from int64 to uint32
func IntToUint32(value int64) (uint32, error) {

	result := uint32(value)

	if int64(result) != value || value < 0 {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToUint64 -- Converts a value from int64 to uint64
func IntToUint64(value int64) (uint64, error) {

	result := uint64(value)

	if int64(result) != value || value < 0 {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// IntToString -- Converts a value from int64 to string
func IntToString(value int64) (string, error) {

	return strconv.FormatInt(value, 10), nil
}
