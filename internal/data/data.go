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

package data

type SyllableType int

const (
	Undefined SyllableType = iota
	Consonant
	Vowel
)

type SungType int

const (
	ChosungType SungType = iota
	JungsungType
	JongsungType
)

const (
	StartHangeul = rune(0xAC00) // '가'
	EndHangeul   = rune(0xD7A3) // '힣'
)

var (
	Consonants = []rune{
		'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ', 'ㄳ', 'ㄵ', 'ㄶ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅄ',
	}
	ChoSung = []rune{
		'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ',
	}
	ChosungComposer = map[string]string{
		"ㄱ":  "ㄱ",
		"ㄲ":  "ㄲ",
		"ㄱㄱ": "ㄲ",
		"ㄴ":  "ㄴ",
		"ㄷ":  "ㄷ",
		"ㄸ":  "ㄸ",
		"ㄷㄷ": "ㄸ",
		"ㄹ":  "ㄹ",
		"ㅁ":  "ㅁ",
		"ㅂ":  "ㅂ",
		"ㅃ":  "ㅃ",
		"ㅂㅂ": "ㅃ",
		"ㅅ":  "ㅅ",
		"ㅆ":  "ㅆ",
		"ㅅㅅ": "ㅆ",
		"ㅇ":  "ㅇ",
		"ㅈ":  "ㅈ",
		"ㅉ":  "ㅉ",
		"ㅈㅈ": "ㅉ",
		"ㅊ":  "ㅊ",
		"ㅋ":  "ㅋ",
		"ㅌ":  "ㅌ",
		"ㅍ":  "ㅍ",
		"ㅎ":  "ㅎ",
	}
	JungSung = []rune{
		'ㅏ', 'ㅐ', 'ㅑ', 'ㅒ', 'ㅓ', 'ㅔ', 'ㅕ', 'ㅖ', 'ㅗ', 'ㅘ', 'ㅙ', 'ㅚ', 'ㅛ', 'ㅜ', 'ㅝ', 'ㅞ', 'ㅟ', 'ㅠ', 'ㅡ', 'ㅢ', 'ㅣ',
	}
	JungsungComposer = map[string]string{
		"ㅏ":   "ㅏ",
		"ㅐ":   "ㅐ",
		"ㅏㅣ":  "ㅐ",
		"ㅑ":   "ㅑ",
		"ㅒ":   "ㅒ",
		"ㅑㅣ":  "ㅒ",
		"ㅐㅐ":  "ㅒ",
		"ㅓ":   "ㅓ",
		"ㅔ":   "ㅔ",
		"ㅓㅣ":  "ㅔ",
		"ㅕ":   "ㅕ",
		"ㅖ":   "ㅖ",
		"ㅕㅣ":  "ㅖ",
		"ㅔㅔ":  "ㅖ",
		"ㅗ":   "ㅗ",
		"ㅘ":   "ㅘ",
		"ㅗㅏ":  "ㅘ",
		"ㅙ":   "ㅙ",
		"ㅗㅐ":  "ㅙ",
		"ㅗㅏㅣ": "ㅙ",
		"ㅚ":   "ㅚ",
		"ㅗㅣ":  "ㅚ",
		"ㅛ":   "ㅛ",
		"ㅜ":   "ㅜ",
		"ㅝ":   "ㅝ",
		"ㅜㅓ":  "ㅝ",
		"ㅞ":   "ㅞ",
		"ㅜㅔ":  "ㅞ",
		"ㅜㅓㅣ": "ㅞ",
		"ㅟ":   "ㅟ",
		"ㅜㅣ":  "ㅟ",
		"ㅠ":   "ㅠ",
		"ㅡ":   "ㅡ",
		"ㅢ":   "ㅢ",
		"ㅡㅣ":  "ㅢ",
		"ㅣ":   "ㅣ",
	}
	JongSung = []rune{
		0, 'ㄱ', 'ㄲ', 'ㄳ', 'ㄴ', 'ㄵ', 'ㄶ', 'ㄷ', 'ㄹ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅁ', 'ㅂ', 'ㅄ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ',
	}
	JongsungComposer = map[string]string{
		"ㄱ": "ㄱ",
		"ㄲ": "ㄲ",
		"ㄳ": "ㄳ",
		"ㄴ": "ㄴ",
		"ㄵ": "ㄵ",
		"ㄶ": "ㄶ",
		"ㄷ": "ㄷ",
		"ㄹ": "ㄹ",
		"ㄺ": "ㄺ",
		"ㄻ": "ㄻ",
		"ㄼ": "ㄼ",
		"ㄽ": "ㄽ",
		"ㄾ": "ㄾ",
		"ㄿ": "ㄿ",
		"ㅀ": "ㅀ",
		"ㅁ": "ㅁ",
		"ㅂ": "ㅂ",
		"ㅄ": "ㅄ",
		"ㅅ": "ㅅ",
		"ㅆ": "ㅆ",
		"ㅇ": "ㅇ",
		"ㅈ": "ㅈ",
		"ㅊ": "ㅊ",
		"ㅋ": "ㅋ",
		"ㅌ": "ㅌ",
		"ㅍ": "ㅍ",
		"ㅎ": "ㅎ",

		"ㄱㄱ": "ㄲ",
		"ㄱㅅ": "ㄳ",
		"ㄴㅈ": "ㄵ",
		"ㄴㅎ": "ㄶ",
		"ㄹㄱ": "ㄺ",
		"ㄹㅁ": "ㄻ",
		"ㄹㅂ": "ㄼ",
		"ㄹㅅ": "ㄽ",
		"ㄹㅌ": "ㄾ",
		"ㄹㅍ": "ㄿ",
		"ㄹㅎ": "ㅀ",
		"ㅂㅅ": "ㅄ",
		"ㅅㅅ": "ㅆ",
	}
)

var ( // QWERTY Keyboard
	// Qwerty keyboard form without shift
	Qwerty = []string{"ㅘ", "ㅙ", "ㅚ", "ㅝ", "ㅞ", "ㅟ", "ㅢ", "ㄳ", "ㄵ", "ㄶ", "ㄺ", "ㄻ", "ㄼ", "ㄽ", "ㄾ", "ㄿ", "ㅀ", "ㅄ"}
	// Qwerty keyboard form without shift
	/*
		fmt.Println(QwertyDecomposer["ㅘ"]) // "ㅗㅏ"
		fmt.Println(QwertyDecomposer["ㄳ"]) // "ㄱㅅ"
		fmt.Println(QwertyDecomposer["ㄲ"]) // "ㄲ"
		// to make "ㄱㄱ" use QwertyOnlyShiftDecomposer1 or QwertyOnlyShiftDecomposer2
	*/
	QwertyDecomposer = map[string]string{
		"ㅘ": "ㅗㅏ",
		"ㅙ": "ㅗㅐ",
		"ㅚ": "ㅗㅣ",
		"ㅝ": "ㅜㅓ",
		"ㅞ": "ㅜㅔ",
		"ㅟ": "ㅜㅣ",
		"ㅢ": "ㅡㅣ",
		"ㄳ": "ㄱㅅ",
		"ㄵ": "ㄴㅈ",
		"ㄶ": "ㄴㅎ",
		"ㄺ": "ㄹㄱ",
		"ㄻ": "ㄹㅁ",
		"ㄼ": "ㄹㅂ",
		"ㄽ": "ㄹㅅ",
		"ㄾ": "ㄹㅌ",
		"ㄿ": "ㄹㅍ",
		"ㅀ": "ㄹㅎ",
		"ㅄ": "ㅂㅅ",
	}
	// Qwerty keyboard form only with shift
	QwertyOnlyShift = []string{"ㄲ", "ㄸ", "ㅃ", "ㅆ", "ㅉ", "ㅒ", "ㅖ"}
	// Qwerty keyboard form only with shift
	/*
		fmt.Println(QwertyOnlyShiftDecomposer1["ㄲ"]) // "ㄱㄱ"
		fmt.Println(QwertyOnlyShiftDecomposer1["ㅒ"]) // "ㅑㅣ"
		fmt.Println(QwertyOnlyShiftDecomposer1["ㅖ"]) // "ㅕㅣ"
	*/
	QwertyOnlyShiftDecomposer1 = map[string]string{
		"ㄲ": "ㄱㄱ",
		"ㄸ": "ㄷㄷ",
		"ㅃ": "ㅂㅂ",
		"ㅆ": "ㅅㅅ",
		"ㅉ": "ㅈㅈ",
		"ㅒ": "ㅑㅣ",
		"ㅖ": "ㅕㅣ",
	}
	// Qwerty keyboard form only with shift
	/*
		fmt.Println(QwertyOnlyShiftDecomposer2["ㄲ"]) // "ㄱㄱ"
		fmt.Println(QwertyOnlyShiftDecomposer2["ㅒ"]) // "ㅐㅐ"
		fmt.Println(QwertyOnlyShiftDecomposer2["ㅖ"]) // "ㅔㅔ"
	*/
	QwertyOnlyShiftDecomposer2 = map[string]string{
		"ㄲ": "ㄱㄱ",
		"ㄸ": "ㄷㄷ",
		"ㅃ": "ㅂㅂ",
		"ㅆ": "ㅅㅅ",
		"ㅉ": "ㅈㅈ",
		"ㅒ": "ㅐㅐ",
		"ㅖ": "ㅔㅔ",
	}
)

var (
	// tense consonants(된소리 자음)
	TenseConsonants = []string{"ㄲ", "ㄸ", "ㅃ", "ㅆ", "ㅉ"}
	// tense consonants(된소리 자음) decomposer
	TenseConsonantsDecomposer = map[string]string{
		"ㄲ": "ㄱㄱ",
		"ㄸ": "ㄷㄷ",
		"ㅃ": "ㅂㅂ",
		"ㅆ": "ㅅㅅ",
		"ㅉ": "ㅈㅈ",
	}
	// complex consonants(복합 자음)
	ComplexConsonants = []string{"ㄳ", "ㄵ", "ㄶ", "ㄺ", "ㄻ", "ㄼ", "ㄽ", "ㄾ", "ㄿ", "ㅀ", "ㅄ"}
	// complex consonants(복합 자음) decomposer
	ComplexConsonantsDecomposer = map[string]string{
		"ㄳ": "ㄱㅅ",
		"ㄵ": "ㄴㅈ",
		"ㄶ": "ㄴㅎ",
		"ㄺ": "ㄹㄱ",
		"ㄻ": "ㄹㅁ",
		"ㄼ": "ㄹㅂ",
		"ㄽ": "ㄹㅅ",
		"ㄾ": "ㄹㅌ",
		"ㄿ": "ㄹㅍ",
		"ㅀ": "ㄹㅎ",
		"ㅄ": "ㅂㅅ",
	}
	// Dipthong(이중 모음)
	Diphthong = []string{"ㅐ", "ㅒ", "ㅔ", "ㅖ", "ㅘ", "ㅙ", "ㅚ", "ㅝ", "ㅞ", "ㅢ"}
	// Dipthong(이중 모음) decomposer
	DiphthongDecomposer = map[string]string{
		"ㅐ": "ㅏㅣ",
		"ㅒ": "ㅑㅣ",
		"ㅔ": "ㅓㅣ",
		"ㅖ": "ㅕㅣ",
		"ㅘ": "ㅗㅏ",
		"ㅙ": "ㅗㅏㅣ",
		"ㅚ": "ㅗㅣ",
		"ㅝ": "ㅜㅓ",
		"ㅞ": "ㅜㅓㅣ",
		"ㅟ": "ㅜㅣ",
		"ㅢ": "ㅡㅣ",
	}
)
