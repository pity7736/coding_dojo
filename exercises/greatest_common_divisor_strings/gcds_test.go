package greatestcommondivisorstrings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"coding_dojo/exercises/greatest_common_divisor_strings"
)

func TestGCDSShould(t *testing.T) {

	t.Run("return empty string when no divisor exists", func(t *testing.T) {
		cases := []struct{ str1, str2 string }{
			{"LEET", "CODE"},
			{"AAAAAB", "AAA"},
		}
		for _, c := range cases {
			got := greatestcommondivisorstrings.GCDStrings(c.str1, c.str2)
			assert.Equal(t, "", got)
		}
	})

	t.Run("return the string itself when both strings are equal", func(t *testing.T) {
		got := greatestcommondivisorstrings.GCDStrings("ABC", "ABC")

		assert.Equal(t, "ABC", got)
	})

	t.Run("return divisor when str2 divides str1", func(t *testing.T) {
		cases := []struct{ str1, str2, expected string }{
			{"ABCABC", "ABC", "ABC"},
			{"ABC", "ABCABC", "ABC"},
			{"ABABAB", "ABAB", "AB"},
			{"ABABAB", "ABAB", "AB"},
			{"ABAB", "ABABABABAB", "AB"},
			{"TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXX"},
		}
		for _, c := range cases {
			got := greatestcommondivisorstrings.GCDStrings(c.str1, c.str2)
			assert.Equal(t, c.expected, got)
		}
	})

}
