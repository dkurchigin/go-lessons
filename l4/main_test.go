package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleGetWordFrequency(t *testing.T) {
	type kv struct {
		Key   string
		Value int
	}

	var expected_ []kv
	expected_ = append(expected_, kv{"тест", 2}, kv{"просто", 1})
	string_ := "Тест, просто тест"
	result_ := GetWordFrequency(string_, 2)
	assert.EqualValues(t, expected_, result_, "Failed simple getting word freq: %s %s", result_, expected_)
}
