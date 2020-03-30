package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestConvertStringsNotValid(t *testing.T) {
	string_ := "a45e"
	_, err := ConvertStrings(string_)
	if err == nil {
		t.Fail()
	}
}

func TestConvertStringsFirstString(t *testing.T) {
	string_ := "a4bc2d5e"
	str_, _ := ConvertStrings(string_)
	require.Equal(t, str_, "aaaabccddddde", "Can't decode %s: result %s", string_, str_)
}

func TestConvertStringsSecondString(t *testing.T) {
	string_ := "abcd"
	str_, _ := ConvertStrings(string_)
	require.Equal(t, str_, "abcd", "Can't decode %s: result %s", string_, str_)
}
