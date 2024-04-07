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

package jamo

import (
	"slices"

	"github.com/ymw0407/jamo/internal/data"
	JamoError "github.com/ymw0407/jamo/pkg/errors"
	"github.com/ymw0407/jamo/pkg/options"
)

// Decompose Hangeul words into Syllables
/*
	// example
	fmt.Println(DecomposeHangeul("한글 is hangeul!")) // "ㅎㅏㄴㄱㅡㄹ is hangeul!"
*/
//* TODO: Allow to apply serveral option (now only first option can apply)
func DecomposeHangeul(hangeuls string, opts ...options.Options) (decomposedHangeul string) {
	filterFunc := func(syllable string) string {
		return syllable
	}

	if len(opts) > 0 {
		option := opts[0]
		filterFunc = option.GetFilterFunc()
	}

	for _, hangeul := range hangeuls {
		if hangeul < data.StartHangeul || hangeul > data.EndHangeul { // non-hangeul character
			decomposedHangeul += string(hangeul)
			continue
		}

		diff := hangeul - data.StartHangeul

		choIdx := diff / (21 * 28)
		jungIdx := (diff % (21 * 28)) / 28
		jongIdx := diff % 28

		chosung := string(data.ChoSung[choIdx])
		jungsung := string(data.JungSung[jungIdx])
		jongsung := ""
		if jongIdx > 0 {
			jongsung = string(data.JongSung[jongIdx])
		}

		decomposedHangeul += filterFunc(chosung) + filterFunc(jungsung) + filterFunc(jongsung)
	}

	return decomposedHangeul
}

// Compose Syllables into hangeul word(not words) //
/*
	// example
	fmt.Println(DecomposeHangeul("한글 is hangeul!")) // "ㅎㅏㄴㄱㅡㄹ is hangeul!"
*/
//* TODO: Allow to apply serveral option (now only first option can apply)
func ComposeHangeul(syllables string) (composedHangeuls []string, err error) {
	sungType := data.ChosungType
	composedHangeuls = append(composedHangeuls, "")
	runeSyllables := []rune(syllables)

	for i := 0; i < len(runeSyllables); i++ {
		stringSyllable := string(runeSyllables[i])
		runeSyllable := rune(runeSyllables[i])
		syllableType := ClassifySyllables(runeSyllable)

		for ii := i + 1; ii < len(runeSyllables); ii++ {
			continuousSyllable := string(runeSyllables[ii])
			if syllableType == ClassifySyllables(runeSyllables[ii]) {
				stringSyllable += continuousSyllable
			} else {
				i = ii - 1
				break
			}
		}

		switch sungType {
		case data.ChosungType: // first chosung only
			chosung, ok := data.ChosungComposer[stringSyllable]
			if !ok {
				return []string{}, JamoError.ErrImpossibleToCompose
			}
			for ii := 0; ii < len(composedHangeuls); ii++ {
				composedHangeuls[ii] += chosung
			}
			sungType = data.JungsungType

		case data.JungsungType:
			jungsung, ok := data.JungsungComposer[stringSyllable]
			if !ok {
				return []string{}, JamoError.ErrImpossibleToCompose
			}
			if len(runeSyllables) == i+1 { // if last jungsung // ㄱㅣ -> 기
				jungsung += " "
			}
			for ii := 0; ii < len(composedHangeuls); ii++ {
				composedHangeuls[ii] += jungsung
			}
			sungType = data.JongsungType

		case data.JongsungType: // only jongsung // only chosung // jongsung + chosung
			if len(runeSyllables) == i+1 { // if last jongsung // ㄷㅗㄴ -> 돈
				jongsung, ok := data.JongsungComposer[stringSyllable]
				if !ok {
					return []string{}, JamoError.ErrImpossibleToCompose
				}
				for ii := 0; ii < len(composedHangeuls); ii++ {
					composedHangeuls[ii] += jongsung
				}
				break
			}

			runeStringSyllable := []rune(stringSyllable)

			switch len(runeStringSyllable) {
			case 1:
				chosung, ok := data.ChosungComposer[stringSyllable] // 0jonsung + 1chosung // ㄱㅏㅂㅏㅇ -> 가방
				if !ok {
					return []string{}, JamoError.ErrImpossibleToCompose
				}
				for ii := 0; ii < len(composedHangeuls); ii++ {
					composedHangeuls[ii] += " " + chosung
				}

			case 2:
				res, ok := data.ChosungComposer[stringSyllable]
				if !ok { // if !ok -> 1jongsung + 1chosung // ㄱㅏㅂㅇㅗㅅ -> 갑옷
					for ii := 0; ii < len(composedHangeuls); ii++ {
						composedHangeuls[ii] += stringSyllable
					}
				} else { // if ok -> 2chosung / 1jongsung + 1chosung // ㄱㅏㅂㅂㅜ -> 가뿌 / 갑부
					length := len(composedHangeuls)
					for ii := 0; ii < length; ii++ {
						composedHangeuls = append(composedHangeuls, composedHangeuls[ii]+stringSyllable)
						composedHangeuls[ii] += " " + res
					}
				}

			case 3:
				jongsung, jongsung2ok := data.JongsungComposer[string(runeStringSyllable[0:2])]
				chosung, chosung2ok := data.JongsungComposer[string(runeStringSyllable[1:3])]
				length := len(composedHangeuls)
				if jongsung2ok && chosung2ok { // 2jongsung + 1chosung / 1jongsung + 2chosung // ㄱㅏㅂㅅㅅㅏㄴ -> 값산 / 갑싼
					for ii := 0; ii < length; ii++ {
						composedHangeuls = append(composedHangeuls, composedHangeuls[ii]+string(runeStringSyllable[0:1])+chosung)
						composedHangeuls[ii] += jongsung + string(runeStringSyllable[2:3])
					}
				} else if jongsung2ok {
					for ii := 0; ii < length; ii++ {
						composedHangeuls[ii] += jongsung + string(runeStringSyllable[2:3])
					}
				} else if chosung2ok {
					for ii := 0; ii < length; ii++ {
						composedHangeuls[ii] += string(runeStringSyllable[0:1]) + chosung
					}
				} else { // error
					return []string{}, JamoError.ErrImpossibleToCompose
				}

			case 4:
				jongsung, ok := data.JongsungComposer[string(runeStringSyllable[0:2])]
				if !ok {
					return []string{}, JamoError.ErrImpossibleToCompose
				}
				chosung, ok := data.ChosungComposer[string(runeStringSyllable[2:4])]
				if !ok {
					return []string{}, JamoError.ErrImpossibleToCompose
				}
				for ii := 0; ii < len(composedHangeuls); ii++ { // 2jongsung + 2chosung // ㄱㅏㅅㅅㅅㅅㅜ -> 갔쑤
					composedHangeuls[ii] += jongsung + chosung
				}
			default:
				return []string{}, JamoError.ErrImpossibleToCompose
			}

			sungType = data.JungsungType
		}
	}

	for i, composedHangeul := range composedHangeuls {
		composedHangeuls[i] = combineHangulSyllables([]rune(composedHangeul))
	}

	return composedHangeuls, err
}

func ClassifySyllables(syllable rune) data.SyllableType {
	switch {
	case slices.Contains(data.Consonants, syllable):
		return data.Consonant
	case slices.Contains(data.JungSung, syllable):
		return data.Vowel
	default:
		return data.Undefined
	}
}

func combineHangulSyllables(jamos []rune) string {
	result := ""
	for i := 0; i < len(jamos); i += 3 {
		choIndex := indexOfRune(data.ChoSung, jamos[i])
		jungIndex := indexOfRune(data.JungSung, jamos[i+1])
		jongIndex := indexOfRune(data.JongSung, jamos[i+2])

		syllableCode := data.StartHangeul + (rune(choIndex) * 21 * 28) + (rune(jungIndex) * 28) + rune(jongIndex)
		result += string(syllableCode)
	}
	return result
}

func indexOfRune(slice []rune, value rune) int {
	if value == ' ' {
		return 0
	}
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
