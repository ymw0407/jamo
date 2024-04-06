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

package options_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymw0407/jamo/pkg/jamo"
	"github.com/ymw0407/jamo/pkg/options"
)

func TestQwertyOption(t *testing.T) {
	t.Run("꺣 테스트! with qwerty options(none shiftOption)", func(t *testing.T) {
		expected := "ㄲㅒㄱㅅ ㅌㅔㅅㅡㅌㅡ!"
		opts := options.Qwerty().SetShiftOption(options.None)
		res := jamo.DecomposeHangeul("꺣 테스트!", *opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꺣 테스트! with qwerty options(shiftOption1)", func(t *testing.T) {
		expected := "ㄱㄱㅑㅣㄱㅅ ㅌㅔㅅㅡㅌㅡ!"
		opts := options.Qwerty().SetShiftOption(options.QwertyShiftOption1)
		res := jamo.DecomposeHangeul("꺣 테스트!", *opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꺣 테스트! with qwerty default options(shiftOption1)", func(t *testing.T) {
		expected := "ㄱㄱㅑㅣㄱㅅ ㅌㅔㅅㅡㅌㅡ!"
		opts := options.Qwerty().SetShiftOption(options.QwertyShiftOption1)
		res := jamo.DecomposeHangeul("꺣 테스트!", *opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꺣 테스트! with qwerty options(shiftOption2)", func(t *testing.T) {
		expected := "ㄱㄱㅐㅐㄱㅅ ㅌㅔㅅㅡㅌㅡ!"
		opts := options.Qwerty().SetShiftOption(options.QwertyShiftOption2)
		res := jamo.DecomposeHangeul("꺣 테스트!", *opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})
}
