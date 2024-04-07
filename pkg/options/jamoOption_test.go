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

func TestJamoOption(t *testing.T) {
	t.Run("꿳테스트 with jamo options(all false)", func(t *testing.T) {
		expected := "ㄲㅞㄳㅌㅔㅅㅡㅌㅡ"
		opts := options.Jamo().SetComplexConsonants(false).SetTenseConsonants(false).SetDiphthong(false)
		res := jamo.DecomposeHangeul("꿳테스트", opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꿳테스트 with jamo options(complex consonants)", func(t *testing.T) {
		expected := "ㄲㅞㄱㅅㅌㅔㅅㅡㅌㅡ"
		opts := options.Jamo().SetComplexConsonants(true).SetTenseConsonants(false).SetDiphthong(false)
		res := jamo.DecomposeHangeul("꿳테스트", opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꿳테스트 with jamo options(tense consonants)", func(t *testing.T) {
		expected := "ㄱㄱㅞㄳㅌㅔㅅㅡㅌㅡ"
		opts := options.Jamo().SetComplexConsonants(false).SetTenseConsonants(true).SetDiphthong(false)
		res := jamo.DecomposeHangeul("꿳테스트", opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꿳테스트 with jamo options(diphthong)", func(t *testing.T) {
		expected := "ㄲㅜㅓㅣㄳㅌㅓㅣㅅㅡㅌㅡ"
		opts := options.Jamo().SetComplexConsonants(false).SetTenseConsonants(false).SetDiphthong(true)
		res := jamo.DecomposeHangeul("꿳테스트", opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꿳테스트 with jamo options(complex consonants, tense consonants)", func(t *testing.T) {
		expected := "ㄱㄱㅞㄱㅅㅌㅔㅅㅡㅌㅡ"
		opts := options.Jamo().SetComplexConsonants(true).SetTenseConsonants(true).SetDiphthong(false)
		res := jamo.DecomposeHangeul("꿳테스트", opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꿳테스트 with jamo options(complex consonants, tense consonants)", func(t *testing.T) {
		expected := "ㄱㄱㅞㄱㅅㅌㅔㅅㅡㅌㅡ"
		opts := options.Jamo().SetComplexConsonants(true).SetTenseConsonants(true).SetDiphthong(false)
		res := jamo.DecomposeHangeul("꿳테스트", opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꿳테스트 with jamo options(complex consonants, diphthong)", func(t *testing.T) {
		expected := "ㄲㅜㅓㅣㄱㅅㅌㅓㅣㅅㅡㅌㅡ"
		opts := options.Jamo().SetComplexConsonants(true).SetTenseConsonants(false).SetDiphthong(true)
		res := jamo.DecomposeHangeul("꿳테스트", opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꿳테스트 with jamo options(tense consonants, diphthong)", func(t *testing.T) {
		expected := "ㄱㄱㅜㅓㅣㄳㅌㅓㅣㅅㅡㅌㅡ"
		opts := options.Jamo().SetComplexConsonants(false).SetTenseConsonants(true).SetDiphthong(true)
		res := jamo.DecomposeHangeul("꿳테스트", opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("꿳테스트 with jamo options(complex consonants, tense consonants, diphthong)", func(t *testing.T) {
		expected := "ㄱㄱㅜㅓㅣㄱㅅㅌㅓㅣㅅㅡㅌㅡ"
		opts := options.Jamo().SetComplexConsonants(true).SetTenseConsonants(true).SetDiphthong(true)
		res := jamo.DecomposeHangeul("꿳테스트", opts)

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})
}
