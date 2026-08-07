package reversewords_test

import (
	"testing"

	"coding_dojo/exercises/reverse_words"
	"github.com/stretchr/testify/assert"
)

func TestReverseWordsShould(t *testing.T) {

	t.Run("return same word when string has one word with no spaces", func(t *testing.T) {
		got := reversewords.ReverseWords("hello")

		assert.Equal(t, "hello", got)
	})

	t.Run("return same word when string has one word with spaces", func(t *testing.T) {
		got := reversewords.ReverseWords(" hello ")

		assert.Equal(t, "hello", got)
	})

	t.Run("return reversed words", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
			want  string
		}{
			{
				name:  "when input has two words separated with space",
				input: "hello world",
				want:  "world hello",
			},
			{
				name:  "when input has three words",
				input: "hello cruel world",
				want:  "world cruel hello",
			},
			{
				name:  "when input has three words with many spaces",
				input: "  hello   cruel   world  ",
				want:  "world cruel hello",
			},
			{
				name:  "example 1",
				input: "the sky is blue",
				want:  "blue is sky the",
			},
			{
				name:  "example 2",
				input: "  hello world  ",
				want:  "world hello",
			},
			{
				name:  "example 3",
				input: "a good   example",
				want:  "example good a",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := reversewords.ReverseWords(tt.input)

				assert.Equal(t, tt.want, got)
			})
		}
	})
}
