package main

import "testing"
import "github.com/stretchr/testify/assert"

type Person struct {
	name string
}

func TestAlwaysReturn1(t *testing.T) {
	result := AlwaysReturn1()
	if result != 1 {
		t.Error("Result should be 1")
	}
}

func TestAlwaysReturn1WithAssert(t *testing.T) {
	result := AlwaysReturn1()
	assert.Equal(t, 1, result, "Result should be 1")
}
