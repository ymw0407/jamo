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
	"github.com/ymw0407/jamo/internal/data"
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
