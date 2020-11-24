package entity

// Connection entity
type Connection struct {
	Source *Airport
	Target *Airport
	Price  float64
}
