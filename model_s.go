// Copyright 2016 The Godaq Authors. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package godaq

const ModelSId = 2

var adcGainsS = []float32{1, 2, 4, 5, 8, 10, 16, 20}

type ModelS struct {
	HwFeatures
}

func NewModelS() *ModelS {
	nInputs := uint(8)
	nOutputs := uint(1)

	return &ModelS{HwFeatures{
		Name:       "OpenDAQ S",
		NLeds:      1,
		NPIOs:      6,
		NInputs:    nInputs,
		NOutputs:   nOutputs,
		NCalibRegs: nOutputs + 2*nInputs,

		Adc: ADC{Bits: 16, Signed: true, VMin: -12.0, VMax: 12.0, Gains: adcGainsS},
		// The DAC has 12 bits, but the firmware transforms the values
		Dac: DAC{Bits: 16, Signed: true, VMin: 0.0, VMax: 4.096},
	}}
}

func (m *ModelS) GetFeatures() HwFeatures {
	return m.HwFeatures
}

func (m *ModelS) GetCalibIndex(isOutput, diffMode, secondStage bool, n, gainId uint) (uint, error) {
	if isOutput {
		if n < 1 || n > m.NOutputs {
			return 0, ErrInvalidOutput
		}
		return n - 1, nil
	}
	if n < 1 || n > m.NInputs || secondStage {
		return 0, ErrInvalidInput
	}
	if diffMode {
		return m.NOutputs + m.NInputs + n - 1, nil
	}
	return m.NOutputs + n - 1, nil
}

func (m *ModelS) CheckValidInputs(pos, neg uint) error {
	if pos < 1 || pos > m.NInputs {
		return ErrInvalidInput
	}
	if neg > 8 {
		return ErrInvalidInput
	}
	return nil
}

func init() {
	// Register this model
	registerModel(ModelSId, NewModelS())
}
