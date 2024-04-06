/*
Copyright 2024 Yun Minwoo. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	"slices"

	"github.com/ymw0407/golang-jamo/internal/data"
)

// option for decompose Qwerty Keyboard form
type QwertyOption struct {
	shiftOption ShiftOption
}

type ShiftOption int

const (
	// default option // none shift option
	/*
		"ㄳ" -> "ㄱㅅ"
		"ㄲ" -> "ㄲ"
		"ㅒ" -> "ㅒ"
	*/
	None ShiftOption = iota
	// option1 // first shift option for qwerty
	/*
		"ㄳ" -> "ㄱㅅ"
		"ㄲ" -> "ㄱㄱ"
		"ㅒ" -> "ㅑㅣ"
	*/
	QwertyShiftOption1
	// option2 // second shift option for qwerty
	/*
		"ㄳ" -> "ㄱㅅ"
		"ㄲ" -> "ㄱㄱ"
		"ㅒ" -> "ㅐㅐ"
	*/
	QwertyShiftOption2
)

// get QwertyOption that option for decompose qwerty keyboard form // option1 is default option
func Qwerty() *QwertyOption {
	return &QwertyOption{
		shiftOption: QwertyShiftOption1,
	}
}

// Option to also decompose tense consonants(된소리 자음, ㄲㄸㅃㅆㅉ) // default is true
/*
	// example
	opts := options.Qwerty().SetTenseConsonants(true) // ㄲ -> ㄱㄱ
*/
func (q *QwertyOption) SetShiftOption(s ShiftOption) *QwertyOption {
	q.shiftOption = s
	return q
}

// GetDecomposeFilter method
func (q QwertyOption) GetFilterFunc() func(syllable string) string {
	return func(syllable string) string {
		if slices.Contains(data.Qwerty, syllable) {
			return data.QwertyDecomposer[syllable]
		} else if slices.Contains(data.QwertyOnlyShift, syllable) {
			switch q.shiftOption {
			case QwertyShiftOption1:
				return data.QwertyOnlyShiftDecomposer1[syllable]
			case QwertyShiftOption2:
				return data.QwertyOnlyShiftDecomposer2[syllable]
			}
		}

		return syllable
	}
}
