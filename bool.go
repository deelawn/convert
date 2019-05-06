package convert

// BoolToFloat32 -- Converts a value from boolean to float32
func BoolToFloat32(value bool) (float32, error) {

	if value {
		return float32(1), nil
	}

	return float32(0), nil
}

// BoolToFloat64 -- Converts a value from boolean to float64
func BoolToFloat64(value bool) (float64, error) {

	if value {
		return float64(1), nil
	}

	return float64(0), nil
}

// BoolToInt -- Converts a value from boolean to int64
func BoolToInt(value bool) (int64, error) {

	if value {
		return int64(1), nil
	}

	return int64(0), nil
}

// BoolToUint -- Converts a value from boolean to uint64
func BoolToUint(value bool) (uint64, error) {

	if value {
		return uint64(1), nil
	}

	return uint64(0), nil
}

// BoolToString -- Converts a value from boolean to string
func BoolToString(value bool) (string, error) {

	// Returns a 1 or 0 string because SQL can convert this to a boolean or integer value whereas it cannot convert
	// "true" or "false" implicitly
	if value {
		return "1", nil
	}

	return "0", nil
}
