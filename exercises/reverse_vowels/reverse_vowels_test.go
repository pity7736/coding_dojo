package reversevowels_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"coding_dojo/exercises/reverse_vowels"
)

func TestReverseVowelsShould(t *testing.T) {
	t.Run("return same string", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{
				name:  "when it has no vowels",
				input: "qwerty",
			},
			{
				name:  "when it has a single vowel",
				input: "ab",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := reversevowels.ReverseVowels(tt.input)

				assert.Equal(t, tt.input, got)
			})
		}
	})

	t.Run("reverse vowels", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
			want  string
		}{
			{
				name:  "when string start and end with a vowel",
				input: "amo",
				want:  "oma",
			},
			{
				name:  "when vowels are in different positions",
				input: "leetcode",
				want:  "leotcede",
			},
			{
				name:  "when there are consecutive consonants between vowels",
				input: "hello",
				want:  "holle",
			},
			{
				name:  "when string has uppercase vowels",
				input: "IceCreAm",
				want:  "AceCreIm",
			},
			{
				name:  "when vowels have accent mark",
				input: "julián",
				want:  "jáliun",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := reversevowels.ReverseVowels(tt.input)

				assert.Equal(t, tt.want, got)
			})
		}

	})
}
