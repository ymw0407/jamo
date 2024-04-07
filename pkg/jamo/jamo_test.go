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
	"github.com/ymw0407/jamo/pkg/jamo"
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

func TestComposeHangeul(t *testing.T) {
	t.Run("깎가-깍까", func(t *testing.T) {
		expected := []string{"깎가", "깍까"}
		res, _ := jamo.ComposeHangeul("ㄱㄱㅏㄱㄱㄱㅏ")

		assert.Equal(t, expected, res)
	})

	t.Run("까까-깍가", func(t *testing.T) {
		expected := []string{"까까", "깍가"}
		res, _ := jamo.ComposeHangeul("ㄱㄱㅏㄱㄱㅏ")

		assert.Equal(t, expected, res)
	})

	t.Run("갃소-각쏘", func(t *testing.T) {
		expected := []string{"갃소", "각쏘"}
		res, _ := jamo.ComposeHangeul("ㄱㅏㄱㅅㅅㅗ")

		assert.Equal(t, expected, res)
	})

	t.Run("테스트1", func(t *testing.T) {
		expected := []string{"테스트"}
		res, _ := jamo.ComposeHangeul("ㅌㅔㅅㅡㅌㅡ")

		assert.Equal(t, expected, res)
	})

	t.Run("테스트2", func(t *testing.T) {
		expected := []string{"테스트"}
		res, _ := jamo.ComposeHangeul("ㅌㅓㅣㅅㅡㅌㅡ")

		assert.Equal(t, expected, res)
	})

	t.Run("기", func(t *testing.T) {
		expected := []string{"기"}
		res, _ := jamo.ComposeHangeul("ㄱㅣ")

		assert.Equal(t, expected, res)
	})

	t.Run("끼", func(t *testing.T) {
		expected := []string{"끼"}
		res, _ := jamo.ComposeHangeul("ㄲㅣ")

		assert.Equal(t, expected, res)
	})

	t.Run("돈", func(t *testing.T) {
		expected := []string{"돈"}
		res, _ := jamo.ComposeHangeul("ㄷㅗㄴ")

		assert.Equal(t, expected, res)
	})

	t.Run("가방", func(t *testing.T) {
		expected := []string{"가방"}
		res, _ := jamo.ComposeHangeul("ㄱㅏㅂㅏㅇ")

		assert.Equal(t, expected, res)
	})

	t.Run("갑옷", func(t *testing.T) {
		expected := []string{"갑옷"}
		res, _ := jamo.ComposeHangeul("ㄱㅏㅂㅇㅗㅅ")

		assert.Equal(t, expected, res)
	})

	t.Run("갑부-가뿌", func(t *testing.T) {
		expected := []string{"가뿌", "갑부"}
		res, _ := jamo.ComposeHangeul("ㄱㅏㅂㅂㅜ")

		assert.Equal(t, expected, res)
	})

	t.Run("값산-갑싼", func(t *testing.T) {
		expected := []string{"값산", "갑싼"}
		res, _ := jamo.ComposeHangeul("ㄱㅏㅂㅅㅅㅏㄴ")

		assert.Equal(t, expected, res)
	})

	t.Run("갔쑤", func(t *testing.T) {
		expected := []string{"갔쑤"}
		res, _ := jamo.ComposeHangeul("ㄱㅏㅅㅅㅅㅅㅜ")

		assert.Equal(t, expected, res)
	})

	t.Run("err 갔 쑤", func(t *testing.T) {
		_, err := jamo.ComposeHangeul("ㄱㅏㅅㅅ ㅅㅅㅜ")

		assert.Error(t, err, err)
	})

	t.Run("err 가e수", func(t *testing.T) {
		_, err := jamo.ComposeHangeul("ㄱㅏeㅅㅜ")

		assert.Error(t, err, err)
	})

	t.Run("err 가 수", func(t *testing.T) {
		_, err := jamo.ComposeHangeul("ㄱㅏ ㅅㅜ")

		assert.Error(t, err, err)
	})

	t.Run("err 가수", func(t *testing.T) {
		_, err := jamo.ComposeHangeul("가ㅅㅜ")

		assert.Error(t, err, err)
	})

}
