package entity

// Connection entity
type Connection struct {
	Source *Airport
	Target *Airport
	Price  float64
}

// NewConnection create a new connection entity
func NewConnection(source *Airport, target *Airport, price float64) (*Connection, error) {
	c := &Connection{
		Source: source,
		Target: target,
		Price:  price,
	}

	err := c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Validate returns an error with bad data
func (c *Connection) Validate() error {
	if c.Source == nil || c.Target == nil {
		return ErrMissingSourceOrTarget
	}

	if c.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}
