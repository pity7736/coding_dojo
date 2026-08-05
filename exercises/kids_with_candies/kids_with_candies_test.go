package kidswithcandies_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"coding_dojo/exercises/kids_with_candies"
)

func TestKidsWithCandiesShould(t *testing.T) {

	t.Run("return nil when candies is empty", func(t *testing.T) {
		got := kidswithcandies.KidsWithCandies([]int{}, 0)

		assert.Empty(t, got)
	})

	t.Run("return true when there is only one kid", func(t *testing.T) {
		got := kidswithcandies.KidsWithCandies([]int{1}, 0)

		assert.Equal(t, []bool{true}, got)
	})

	t.Run("return false for kids that cannot reach the greatest even with extra candies", func(t *testing.T) {
		cases := []struct {
			candies      []int
			extraCandies int
			expected     []bool
		}{
			{
				[]int{12, 1, 12},
				10,
				[]bool{true, false, true},
			},
			{
				[]int{2, 3, 5, 1, 3},
				3,
				[]bool{true, true, true, false, true},
			},
			{
				[]int{4, 2, 1, 1, 2},
				1,
				[]bool{true, false, false, false, false},
			},
		}
		for _, c := range cases {
			got := kidswithcandies.KidsWithCandies(c.candies, c.extraCandies)
			assert.Equal(t, c.expected, got)
		}

	})

}
