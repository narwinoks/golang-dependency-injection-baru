package test

import (
	"golang-rest-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, clean := simple.InitializedConnection("FILE TESTING")
	assert.NotNil(t, connection)
	clean()
}
