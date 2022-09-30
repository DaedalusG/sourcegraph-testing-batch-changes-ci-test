package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, "foo", "faz")
}

func Test2(t *testing.T) {
	assert.Equal(t, "bar", "bar")
}
