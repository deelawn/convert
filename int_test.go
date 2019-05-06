package convert

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IntToBool(t *testing.T) {

	trueBool, trueErr := IntToBool(int64(1))
	falseBool, falseErr := IntToBool(int64(0))
	_, err := IntToBool(int64(100))

	assert.True(t, trueBool)
	assert.NoError(t, trueErr)
	assert.False(t, falseBool)
	assert.NoError(t, falseErr)
	assert.Error(t, err)
}

func Test_IntToFloat32(t *testing.T) {

	okResult, okErr := IntToFloat32(int64(10))
	_, err := IntToFloat32(int64(float32MaxPreciseValue + 1))

	assert.EqualValues(t, 10, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_IntToFloat64(t *testing.T) {

	okResult, okErr := IntToFloat64(int64(10))
	_, err := IntToFloat64(int64(float64MaxPreciseValue + 1))

	assert.EqualValues(t, 10, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_IntToDefaultInt(t *testing.T) {

	okResult, okErr := IntToDefaultInt(int64(19))
	_, err := IntToDefaultInt(int64(math.MaxInt64))

	assert.EqualValues(t, 19, okResult)
	assert.NoError(t, okErr)

	if strconv.IntSize < 64 {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}
}

func Test_IntToInt8(t *testing.T) {

	okResult, okErr := IntToInt8(int64(math.MaxInt8))
	_, err := IntToInt8(int64(math.MaxInt8 + 1))

	assert.EqualValues(t, math.MaxInt8, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_IntToInt16(t *testing.T) {

	okResult, okErr := IntToInt16(int64(math.MaxInt16))
	_, err := IntToInt16(int64(math.MaxInt16 + 1))

	assert.EqualValues(t, math.MaxInt16, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_IntToInt32(t *testing.T) {

	okResult, okErr := IntToInt32(int64(math.MaxInt32))
	_, err := IntToInt32(int64(math.MaxInt32 + 1))

	assert.EqualValues(t, math.MaxInt32, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_IntToUint(t *testing.T) {

	okResult, okErr := IntToUint(int64(19))
	_, err := IntToUint(int64(-60))

	assert.EqualValues(t, 19, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_IntToUint8(t *testing.T) {

	okResult, okErr := IntToUint8(int64(math.MaxUint8))
	_, err := IntToUint8(int64(math.MaxUint8 + 1))

	assert.EqualValues(t, math.MaxUint8, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_IntToUint16(t *testing.T) {

	okResult, okErr := IntToUint16(int64(math.MaxUint16))
	_, err := IntToUint16(int64(math.MaxUint16 + 1))

	assert.EqualValues(t, math.MaxUint16, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_IntToUint32(t *testing.T) {

	okResult, okErr := IntToUint32(int64(9000))
	_, err := IntToUint32(int64(-80))

	assert.EqualValues(t, 9000, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_IntToUint64(t *testing.T) {

	okResult, okErr := IntToUint64(int64(123456789))
	_, err := IntToUint64(int64(-887))

	assert.EqualValues(t, 123456789, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_IntToString(t *testing.T) {

	okResult, okErr := IntToString(int64(math.MaxInt64))

	assert.Equal(t, "9223372036854775807", okResult)
	assert.NoError(t, okErr)
}
