package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringToBool(t *testing.T) {

	trueBool, trueErr := StringToBool("T")
	falseBool, falseErr := StringToBool("0")
	_, err := StringToBool("X")

	assert.True(t, trueBool)
	assert.NoError(t, trueErr)
	assert.False(t, falseBool)
	assert.NoError(t, falseErr)
	assert.Error(t, err)
}

func Test_StringToFloat32(t *testing.T) {

	okResult, okErr := StringToFloat32("27.43")
	_, err := StringToFloat32("asdf")

	assert.Equal(t, float32(27.43), okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToFloat64(t *testing.T) {

	okResult, okErr := StringToFloat64("27.43")
	_, err := StringToFloat64("asdf")

	assert.Equal(t, float64(27.43), okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToInt(t *testing.T) {

	okResult, okErr := StringToInt("19 ")
	_, err := StringToInt("asdf")

	assert.EqualValues(t, 19, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToInt8(t *testing.T) {

	okResult, okErr := StringToInt8(" 19")
	_, err := StringToInt8("256")

	assert.EqualValues(t, 19, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToInt16(t *testing.T) {

	okResult, okErr := StringToInt16("400")
	_, err := StringToInt16("65536")

	assert.EqualValues(t, 400, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToInt32(t *testing.T) {

	okResult, okErr := StringToInt32("400")
	_, err := StringToInt32("4294967296")

	assert.EqualValues(t, 400, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToInt64(t *testing.T) {

	okResult, okErr := StringToInt64("400")
	_, err := StringToInt64("asdf")

	assert.EqualValues(t, 400, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToUint(t *testing.T) {

	okResult, okErr := StringToUint("19 ")
	_, err := StringToUint("asdf")

	assert.EqualValues(t, 19, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToUint8(t *testing.T) {

	okResult, okErr := StringToUint8(" 19")
	_, err := StringToUint8("512")

	assert.EqualValues(t, 19, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToUint16(t *testing.T) {

	okResult, okErr := StringToUint16("400")
	_, err := StringToUint16("131072")

	assert.EqualValues(t, 400, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToUint32(t *testing.T) {

	okResult, okErr := StringToUint32("400")
	_, err := StringToUint32(" 8589934592")

	assert.EqualValues(t, 400, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}

func Test_StringToUint64(t *testing.T) {

	okResult, okErr := StringToUint64("400")
	_, err := StringToUint64("asdf")

	assert.EqualValues(t, 400, okResult)
	assert.NoError(t, okErr)
	assert.Error(t, err)
}
