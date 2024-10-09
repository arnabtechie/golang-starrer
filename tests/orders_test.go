package tests

import (
	"testing"

	"github.com/arnabtechie/golang-starrer/controllers/orders"
	"github.com/stretchr/testify/assert"
)

func TestOrderAdd(t *testing.T) {
	result := orders.Add(5, 3)
	assert.Equal(t, 8, result, "they should be equal")
}

func TestOrderSubtract(t *testing.T) {
	result := orders.Subtract(5, 3)
	assert.Equal(t, 2, result, "they should be equal")
}
