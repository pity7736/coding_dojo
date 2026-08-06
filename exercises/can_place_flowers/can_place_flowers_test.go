package canplaceflowers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"coding_dojo/exercises/can_place_flowers"
)

func TestCanPlaceFlowersShould(t *testing.T) {

	t.Run("return true when", func(t *testing.T) {
		tests := []struct {
			name      string
			flowerbed []int
			n         int
		}{
			{
				name:      "flowerbed is empty and n is zero",
				flowerbed: []int{},
				n:         0,
			},
			{
				name:      "flowerbed has values and n is one",
				flowerbed: []int{0, 0, 0},
				n:         1,
			},
			{
				name:      "flowerbed has one spot and n is one",
				flowerbed: []int{0},
				n:         1,
			},
			{
				name:      "flowerbed has available spot at the beginning",
				flowerbed: []int{0, 0, 1},
				n:         1,
			},
			{
				name:      "two flowers can be planted in all empty flowerbed",
				flowerbed: []int{0, 0, 0},
				n:         2,
			},
			{
				name:      "flowers can be planted after an existing flower at the start",
				flowerbed: []int{0, 0, 0, 0, 0},
				n:         3,
			},
			{
				name:      "example 1: one flower between two existing",
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := canplaceflowers.CanPlaceFlowers(tt.flowerbed, tt.n)
				assert.True(t, got)

			})
		}
	})

	t.Run("return false when", func(t *testing.T) {
		tests := []struct {
			name      string
			flowerbed []int
			n         int
		}{
			{
				name:      "flowerbed is empty and n is one",
				flowerbed: []int{},
				n:         1,
			},
			{
				name:      "flowerbed has not spot and n is one",
				flowerbed: []int{0, 1, 0},
				n:         1,
			},
			{
				name:      "flowerbed has one spot and n is two",
				flowerbed: []int{0},
				n:         2,
			},
			{
				name:      "flowerbed has adjacent empty spots but only one can be planted",
				flowerbed: []int{0, 0, 0, 0, 1},
				n:         3,
			},
			{
				name:      "example 2: two flowers cannot fit between two existing",
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := canplaceflowers.CanPlaceFlowers(tt.flowerbed, tt.n)
				assert.False(t, got)

			})
		}
	})
}
