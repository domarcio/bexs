package entity

import (
	"regexp"
	"strings"
	"unicode"
)

// Airport entity
type Airport struct {
	Code string
}

// NewAirport create a new airpot entity
func NewAirport(code string) (*Airport, error) {
	code = strings.TrimSpace(code)
	c := &Airport{
		Code: code,
	}

	err := c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Validate returns an error with bad data
func (c *Airport) Validate() error {
	if c.Code == "" {
		return ErrEmptyCode
	}

	lenCode := len(c.Code)
	if lenCode > 3 || lenCode < 3 {
		return ErrLenCode
	}

	rg, err := regexp.Match("^[0-9]", []byte(c.Code))
	if err != nil {
		return err
	}
	if rg {
		return ErrNumberCode
	}

	for _, s := range c.Code {
		if !unicode.IsUpper(s) {
			return ErrInvalidCaseCode
		}
	}

	return nil
}
