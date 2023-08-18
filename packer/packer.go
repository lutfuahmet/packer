package packer

// Packer defines the interface for calculating packs.
type Packer interface {
	CalculatePacks(itemsQuantity uint) map[uint]uint
	UpdatePackSizes(packSizes []uint)
}
