package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BoolToFloat32(t *testing.T) {

	trueResult, trueErr := BoolToFloat32(true)
	falseResult, falseErr := BoolToFloat32(false)

	assert.EqualValues(t, 1, trueResult)
	assert.NoError(t, trueErr)
	assert.EqualValues(t, 0, falseResult)
	assert.NoError(t, falseErr)
}

func Test_BoolToFloat64(t *testing.T) {

	trueResult, trueErr := BoolToFloat64(true)
	falseResult, falseErr := BoolToFloat64(false)

	assert.EqualValues(t, 1, trueResult)
	assert.NoError(t, trueErr)
	assert.EqualValues(t, 0, falseResult)
	assert.NoError(t, falseErr)
}

func Test_BoolToInt(t *testing.T) {

	trueResult, trueErr := BoolToInt(true)
	falseResult, falseErr := BoolToInt(false)

	assert.EqualValues(t, 1, trueResult)
	assert.NoError(t, trueErr)
	assert.EqualValues(t, 0, falseResult)
	assert.NoError(t, falseErr)
}

func Test_BoolToUint(t *testing.T) {

	trueResult, trueErr := BoolToUint(true)
	falseResult, falseErr := BoolToUint(false)

	assert.EqualValues(t, 1, trueResult)
	assert.NoError(t, trueErr)
	assert.EqualValues(t, 0, falseResult)
	assert.NoError(t, falseErr)
}

func Test_BoolToString(t *testing.T) {

	trueResult, trueErr := BoolToString(true)
	falseResult, falseErr := BoolToString(false)

	assert.Equal(t, "1", trueResult)
	assert.NoError(t, trueErr)
	assert.Equal(t, "0", falseResult)
	assert.NoError(t, falseErr)
}
