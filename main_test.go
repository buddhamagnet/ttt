package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInvalidMove(t *testing.T) {
	player := Player{}
	_, err := getMove(player)
	assert.Error(t, err)

}
