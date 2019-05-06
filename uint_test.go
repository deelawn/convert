package convert

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UintToBool(t *testing.T) {

	trueBool, trueErr := UintToBool(uint64(1))
	falseBool, falseErr := UintToBool(uint64(0))
	_, err := UintToBool(uint64(100))

	assert.True(t, trueBool)
	assert.NoError(t, trueErr)
	assert.False(t, falseBool)
	assert.NoError(t, falseErr)
	assert.Error(t, err)
}

func Test_UintToFloat32(t *testing.T) {

	okResult, okErr := UintToFloat32(uint64(10))
	_, err := UintToFloat32(uint64(float32MaxPreciseValue + 1))

	assert.EqualValues(t, 10, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_UintToFloat64(t *testing.T) {

	okResult, okErr := UintToFloat64(uint64(10))
	_, err := UintToFloat64(uint64(float64MaxPreciseValue + 1))

	assert.EqualValues(t, 10, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_UintToInt(t *testing.T) {

	okResult, okErr := UintToInt(uint64(19))
	_, err := UintToInt(uint64(math.MaxUint64))

	assert.EqualValues(t, 19, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_UintToInt8(t *testing.T) {

	okResult, okErr := UintToInt8(uint64(math.MaxInt8))
	_, err := UintToInt8(uint64(math.MaxInt8 + 1))

	assert.EqualValues(t, math.MaxInt8, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_UintToInt16(t *testing.T) {

	okResult, okErr := UintToInt16(uint64(math.MaxInt16))
	_, err := UintToInt16(uint64(math.MaxInt16 + 1))

	assert.EqualValues(t, math.MaxInt16, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_UintToInt32(t *testing.T) {

	okResult, okErr := UintToInt32(uint64(math.MaxInt32))
	_, err := UintToInt32(uint64(math.MaxInt32 + 1))

	assert.EqualValues(t, math.MaxInt32, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_UintToInt64(t *testing.T) {

	okResult, okErr := UintToInt64(uint64(math.MaxInt64))
	_, err := UintToInt64(uint64(math.MaxInt64 + 1))

	assert.EqualValues(t, math.MaxInt64, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_UintToDefaultUint(t *testing.T) {

	okResult, okErr := UintToDefaultUint(uint64(19))
	_, err := UintToDefaultUint(uint64(math.MaxUint64))

	assert.EqualValues(t, 19, okResult)
	assert.NoError(t, okErr)

	if strconv.IntSize == 64 {
		assert.NoError(t, err)
	} else {
		assert.Error(t, err)
	}
}

func Test_UintToUint8(t *testing.T) {

	okResult, okErr := UintToUint8(uint64(math.MaxUint8))
	_, err := UintToUint8(uint64(math.MaxUint8 + 1))

	assert.EqualValues(t, math.MaxUint8, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_UintToUint16(t *testing.T) {

	okResult, okErr := UintToUint16(uint64(math.MaxUint16))
	_, err := UintToUint16(uint64(math.MaxUint16 + 1))

	assert.EqualValues(t, math.MaxUint16, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_UintToUint32(t *testing.T) {

	okResult, okErr := UintToUint32(math.MaxUint32)
	_, err := UintToUint32(math.MaxUint32 + 1)

	assert.EqualValues(t, math.MaxUint32, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_UintToString(t *testing.T) {

	okResult, okErr := UintToString(uint64(math.MaxInt64))

	assert.Equal(t, "9223372036854775807", okResult)
	assert.NoError(t, okErr)
}
