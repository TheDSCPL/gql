package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gituhub.com/TheDSCPL/gql/graph/model"
)

func TestInvalidID(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		assert.True(t, model.NewID().IsValid())
	}

	assert.False(t, model.ID("invalid").IsValid())
	// UUID v1 is not acceptable
	assert.False(t, model.ID("5b301e64-a39b-11ed-a8fc-0242ac120002").IsValid())
}
