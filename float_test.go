package convert

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FloatToBool(t *testing.T) {

	trueBool, trueErr := FloatToBool(float64(1))
	falseBool, falseErr := FloatToBool(float64(0))
	_, err := FloatToBool(float64(100.2))

	assert.True(t, trueBool)
	assert.NoError(t, trueErr)
	assert.False(t, falseBool)
	assert.NoError(t, falseErr)
	assert.Error(t, err)
}

func Test_FloatToFloat32(t *testing.T) {

	okResult, okErr := FloatToFloat32(float64(99.99))
	_, parseErr := FloatToFloat32(math.MaxFloat64)
	_, err := FloatToFloat32(float64(25.9440404059))

	assert.EqualValues(t, 99.99, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, parseErr)
	assert.Error(t, err)
}

func Test_FloatToInt(t *testing.T) {

	okResult, okErr := FloatToInt(float64(19))
	_, err := FloatToInt(float64(60.32))

	assert.EqualValues(t, 19, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_FloatToInt8(t *testing.T) {

	okResult, okErr := FloatToInt8(float64(math.MaxInt8))
	_, err := FloatToInt8(float64(math.MaxInt8 + 1))

	assert.EqualValues(t, math.MaxInt8, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_FloatToInt16(t *testing.T) {

	okResult, okErr := FloatToInt16(float64(math.MaxInt16))
	_, err := FloatToInt16(float64(math.MaxInt16 + 1))

	assert.EqualValues(t, math.MaxInt16, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_FloatToInt32(t *testing.T) {

	okResult, okErr := FloatToInt32(float64(math.MaxInt32))
	_, err := FloatToInt32(float64(math.MaxInt32 + 1))

	assert.EqualValues(t, math.MaxInt32, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_FloatToInt64(t *testing.T) {

	okResult, okErr := FloatToInt64(float64(123456789))
	_, err := FloatToInt64(float64(89484.34834))

	assert.EqualValues(t, 123456789, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_FloatToUint(t *testing.T) {

	okResult, okErr := FloatToUint(float64(19))
	_, err := FloatToUint(float64(60.32))

	assert.EqualValues(t, 19, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_FloatToUint8(t *testing.T) {

	okResult, okErr := FloatToUint8(float64(math.MaxUint8))
	_, err := FloatToUint8(float64(math.MaxUint8 + 1))

	assert.EqualValues(t, math.MaxUint8, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_FloatToUint16(t *testing.T) {

	okResult, okErr := FloatToUint16(float64(math.MaxUint16))
	_, err := FloatToUint16(float64(math.MaxUint16 + 1))

	assert.EqualValues(t, math.MaxUint16, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_FloatToUint32(t *testing.T) {

	okResult, okErr := FloatToUint32(float64(9000))
	_, err := FloatToUint32(float64(9000.1))

	assert.EqualValues(t, 9000, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_FloatToUint64(t *testing.T) {

	okResult, okErr := FloatToUint64(float64(123456789))
	_, err := FloatToUint64(float64(89484.34834))

	assert.EqualValues(t, 123456789, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_FloatToString(t *testing.T) {

	okResult, okErr := FloatToString(float64(float64MaxPreciseValue))

	assert.Equal(t, "9007199254740992", okResult)
	assert.NoError(t, okErr)
}
