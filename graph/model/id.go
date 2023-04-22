package model

import (
	"regexp"

	"github.com/google/uuid"
)

func init() {
	uuid.EnableRandPool()
}

func NewID() ID {
	return ID(uuid.New().String())
}

// ID is a type alias for a string that represents a UUID v4
type ID string

var uuidV4Regex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

func (id ID) IsValid() bool {
	return uuidV4Regex.MatchString(string(id))
}
