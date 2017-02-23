package godaq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTP04ABCalibIndex(t *testing.T) {
	hw := newModelTP04AB()
	assert.Equal(t, "TP04AB", hw.Name)
	assert.EqualValues(t, 10, hw.NCalibRegs)

	for i := uint(0); i < hw.NOutputs; i++ {
		idx, err := hw.GetCalibIndex(true, false, false, i+1, 0)
		assert.NoError(t, err)
		assert.EqualValues(t, i, idx)
	}

	for i := uint(0); i < hw.NInputs; i++ {
		// first stage calibration slots
		idx, err := hw.GetCalibIndex(false, false, false, i+1, 0)
		assert.EqualValues(t, hw.NOutputs+i, idx)
		assert.Nil(t, err)
		// second stage calibration slots
		idx, err = hw.GetCalibIndex(false, false, true, i+1, 0)
		assert.EqualValues(t, hw.NOutputs+hw.NInputs+i, idx)
		assert.Nil(t, err)
	}
}