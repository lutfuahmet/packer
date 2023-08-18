package packer

import (
	"math"
	"packer/utils"
	"sort"
)

// DefaultPackageCalculator implements Packer.
type DefaultPackageCalculator struct {
	packSizes []uint
}

// NewDefaultPackageCalculator creates a new DefaultPackageCalculator
func NewDefaultPackageCalculator(packSizes []uint) *DefaultPackageCalculator {
	return &DefaultPackageCalculator{packSizes: packSizes}
}

// UpdatePackSizes updates pack sizes.
func (p *DefaultPackageCalculator) UpdatePackSizes(packSizes []uint) {
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] < packSizes[j]
	})
	p.packSizes = packSizes
}

// CalculatePacks calculates the required number of packs for a given item quantity.
func (p *DefaultPackageCalculator) CalculatePacks(itemQuantity uint) map[uint]uint {
	return CalculatePacksWithPackSizes(p.packSizes, itemQuantity)
}

// CalculatePacksWithPackSizes calculates the required number of packs for a given item quantity and packSizes
func CalculatePacksWithPackSizes(packSizes []uint, itemQuantity uint) map[uint]uint {
	if len(packSizes) == 0 {
		return nil
	}
	packCounts := make(map[uint]uint)

	smallest := packSizes[0]

	for i := len(packSizes) - 1; i >= 0; i-- {
		packSize := packSizes[i]
		packCount := itemQuantity / packSize
		itemQuantity %= packSize
		if packCount > 0 {
			packCounts[packSize] = packCount
		}
	}
	if len(packCounts) == 0 {
		// added smallest package
		packCounts[smallest] = 1
		itemQuantity = 0
	}
	// check remaining items
	if itemQuantity > 0 {
		// smallest
		bestSize := smallest
		minDifference := math.MaxInt
		for _, size := range packSizes {
			difference := size - itemQuantity
			if difference >= 0 && int(difference) < minDifference {
				bestSize = size
				minDifference = int(difference)
			}
		}

		// optimize pack sizes

		firstBreakPointIndex := utils.FindFirstBreakPoint(packSizes)
		var firstBreakPoint uint
		if firstBreakPointIndex > -1 {
			firstBreakPoint = packSizes[firstBreakPointIndex]
		} else {
			firstBreakPoint = bestSize
		}

		if bestSize != firstBreakPoint && firstBreakPointIndex > 0 {
			maxLess := packSizes[firstBreakPointIndex-1] // largest number smaller than first break point
			if itemQuantity <= (firstBreakPoint - maxLess) {
				packCounts[firstBreakPoint]++
				packCounts[maxLess]--
			}
		} else {
			// Update pack counts
			packCounts[bestSize]++
		}

		for _, size := range packSizes {
			count := packCounts[size]
			if count == 0 {
				continue
			}
			itemCount := count * size
			if count > 1 && utils.Contains(packSizes, itemCount) {
				delete(packCounts, size)
				packCounts[itemCount]++
			}
		}
		newPackCount := map[uint]uint{}
		// delete 0 values
		for size, count := range packCounts {
			if count > 0 {
				newPackCount[size] = count
			}
		}
		packCounts = newPackCount
	}

	return packCounts
}
