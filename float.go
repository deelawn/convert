package convert

import (
	"errors"
	"strconv"
)

// FloatToBool -- Converts a value from float64 to boolean
func FloatToBool(value float64) (bool, error) {

	if float64(1) == value {
		return true, nil
	} else if float64(0) == value {
		return false, nil
	}

	return false, errors.New(cannotConvertErrMsg)
}

// FloatToFloat32 -- Converts a value from float64 to float32
func FloatToFloat32(value float64) (float32, error) {

	valueStr, _ := FloatToString(value)
	result64, parseErr := strconv.ParseFloat(valueStr, 32)
	result := float32(result64)

	if parseErr != nil {
		return result, parseErr
	}

	resultStr := strconv.FormatFloat(float64(result), 'f', -1, 32)

	if resultStr != valueStr {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToInt -- Converts a value from float64 to int
func FloatToInt(value float64) (int, error) {

	result := int(value)

	if float64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToInt8 -- Converts a value from float64 to int8
func FloatToInt8(value float64) (int8, error) {

	result := int8(value)

	if float64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToInt16 -- Converts a value from float64 to int16
func FloatToInt16(value float64) (int16, error) {

	result := int16(value)

	if float64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToInt32 -- Converts a value from float64 to int32
func FloatToInt32(value float64) (int32, error) {

	result := int32(value)

	if float64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToInt64 -- Converts a value from float64 to int64
func FloatToInt64(value float64) (int64, error) {

	result := int64(value)

	if float64(result) != value {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToUint -- Converts a value from float64 to uint
func FloatToUint(value float64) (uint, error) {

	result := uint(value)

	if float64(result) != value || value < 0 {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToUint8 -- Converts a value from float64 to uint8
func FloatToUint8(value float64) (uint8, error) {

	result := uint8(value)

	if float64(result) != value || value < 0 {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToUint16 -- Converts a value from float64 to uint16
func FloatToUint16(value float64) (uint16, error) {

	result := uint16(value)

	if float64(result) != value || value < 0 {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToUint32 -- Converts a value from float64 to uint32
func FloatToUint32(value float64) (uint32, error) {

	result := uint32(value)

	if float64(result) != value || value < 0 {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToUint64 -- Converts a value from float64 to uint64
func FloatToUint64(value float64) (uint64, error) {

	result := uint64(value)

	if float64(result) != value || value < 0 {
		return result, errors.New(lossOfPrecisionErrorMsg)
	}

	return result, nil
}

// FloatToString -- Converts a value from float64 to string
func FloatToString(value float64) (string, error) {

	return strconv.FormatFloat(value, 'f', -1, 64), nil
}
