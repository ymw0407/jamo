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

// option for decompose tense constants(된소리 자음, ㄲㄸㅃㅆㅉ), complex consonants(복합 자음, ㄳㄵㄶㄺㄻㄼㄽㄾㄿㅀㅄ), dipthongs(이중 모음, ㅐㅔㅘㅙㅚㅝㅞㅢ)
type JamoOption struct {
	tenseConsonants   bool
	complexConsonants bool
	diphthong         bool
}

// get JamoOption that option for decompose tense constants(된소리 자음, ㄲㄸㅃㅆㅉ), complex consonants(복합 자음, ㄳㄵㄶㄺㄻㄼㄽㄾㄿㅀㅄ), dipthongs(이중 모음, ㅐㅔㅘㅙㅚㅝㅞㅢ)
func Jamo() *JamoOption {
	return &JamoOption{
		tenseConsonants:   true,
		complexConsonants: true,
		diphthong:         true,
	}
}

// Option to also decompose tense consonants(된소리 자음, ㄲㄸㅃㅆㅉ) // default is true
/*
	// example
	opts := options.Jamo().SetTenseConsonants(true) // ㄲ -> ㄱㄱ
*/
func (j *JamoOption) SetTenseConsonants(b bool) *JamoOption {
	j.tenseConsonants = b
	return j
}

// Option to also decompose complex consonants(복합 자음, ㄳㄵㄶㄺㄻㄼㄽㄾㄿㅀㅄ) // default is true
/*
	// example
	opts := options.Jamo().SetComplexConsonants(true) // ㄳ -> ㄳ
*/
func (j *JamoOption) SetComplexConsonants(b bool) *JamoOption {
	j.complexConsonants = b
	return j
}

// Option to also decompose dipthongs(이중 모음, ㅐㅒㅔㅖㅘㅙㅚㅝㅞㅢ) // default is true
/*
	// example
	opts := options.Jamo().SetDiphthong(true) // ㅞ -> ㅜㅓㅣ
*/
func (j *JamoOption) SetDiphthong(b bool) *JamoOption {
	j.diphthong = b
	return j
}

// GetDecomposeFilter method
func (j JamoOption) GetFilterFunc() func(syllable string) string {
	return func(syllable string) string {
		if j.tenseConsonants && slices.Contains(data.TenseConsonants, syllable) { // ㄲㄸㅃㅆㅉ
			return data.TenseConsonantsDecomposer[syllable]
		} else if j.complexConsonants && slices.Contains(data.ComplexConsonants, syllable) { // ㄳㄵㄶㄺㄻㄼㄽㄾㄿㅀㅄ
			return data.ComplexConsonantsDecomposer[syllable]
		} else if j.diphthong && slices.Contains(data.Diphthong, syllable) { // ㅐㅒㅔㅖㅘㅙㅚㅝㅞㅢ
			return data.DiphthongDecomposer[syllable]
		}

		return syllable
	}
}
