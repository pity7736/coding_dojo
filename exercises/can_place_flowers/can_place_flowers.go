package canplaceflowers

func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	availableSpots := 0
	plantedFlowerIndex := 0
	for i, spot := range flowerbed {
		if canPlantAt(flowerbed, plantedFlowerIndex, spot, i) {
			availableSpots++
			plantedFlowerIndex = i
		}
		if availableSpots >= n {
			return true
		}
	}
	return false
}

func canPlantAt(flowerbed []int, plantedFlowerIndex, spot, i int) bool {
	return spot == 0 && plantedFlowerIndex != i-1 && isSpotAvailable(flowerbed, i-1) && isSpotAvailable(flowerbed, i+1)
}

func isSpotAvailable(flowerbed []int, n int) bool {
	if n < 0 || n >= len(flowerbed) {
		return true
	} else {
		return flowerbed[n] == 0
	}
}
