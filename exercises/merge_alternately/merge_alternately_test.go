package mergealternately_test

import (
	"coding_dojo/exercises/merge_alternately"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternatelyShould(t *testing.T) {

	t.Run("merge two equals length strings", func(t *testing.T) {
		got := mergealternately.MergeAlternately("abc", "pqr")

		assert.Equal(t, "apbqcr", got)
	})

	t.Run("append remaining chars when word2 is longer", func(t *testing.T) {
		got := mergealternately.MergeAlternately("ab", "pqrs")

		assert.Equal(t, "apbqrs", got)
	})

	t.Run("append remaining chars when word1 is longer", func(t *testing.T) {
		got := mergealternately.MergeAlternately("abcd", "pq")

		assert.Equal(t, "apbqcd", got)
	})

	t.Run("merge two equal length unicode strings", func(t *testing.T) {
		got := mergealternately.MergeAlternately("áéí", "óúñ")

		assert.Equal(t, "áóéúíñ", got)
	})

	t.Run("merge unicode strings when word2 is longer", func(t *testing.T) {
		got := mergealternately.MergeAlternately("日", "本語")

		assert.Equal(t, "日本語", got)
	})

	t.Run("return word2 when word1 is empty", func(t *testing.T) {
		got := mergealternately.MergeAlternately("", "abc")

		assert.Equal(t, "abc", got)
	})

	t.Run("return word1 when word2 is empty", func(t *testing.T) {
		got := mergealternately.MergeAlternately("abc", "")

		assert.Equal(t, "abc", got)
	})

	t.Run("return empty string when both are empty", func(t *testing.T) {
		got := mergealternately.MergeAlternately("", "")

		assert.Equal(t, "", got)
	})
}
