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

package jamo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymw0407/golang-jamo/pkg/jamo"
)

func TestDecomposeHangeul(t *testing.T) {
	t.Run("한글 is hangeul!", func(t *testing.T) {
		expected := "ㅎㅏㄴㄱㅡㄹ is hangeul!"
		res := jamo.DecomposeHangeul("한글 is hangeul!")

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("ㄱ is giyeok!", func(t *testing.T) {
		expected := "ㄱ is giyeok!"
		res := jamo.DecomposeHangeul("ㄱ is giyeok!")

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})

	t.Run("윤민우", func(t *testing.T) {
		expected := "ㅇㅠㄴㅁㅣㄴㅇㅜ"
		res := jamo.DecomposeHangeul("윤민우")

		assert.Equal(t, expected, res, res+" : `"+expected+"` is expected!")
	})
}
