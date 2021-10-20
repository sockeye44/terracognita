package util_test

import (
	"testing"

	"github.com/sockeye44/terracognita/util"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeName(t *testing.T) {
	tests := []struct {
		Name     string
		In       string
		Expected string
	}{
		{
			Name:     "NoChange",
			In:       "in",
			Expected: "in",
		},
		{
			Name:     "UpperCase",
			In:       "IN",
			Expected: "in",
		},
		{
			Name:     "Invalid",
			In:       ":a",
			Expected: "_a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, tt.Expected, util.NormalizeName(tt.In))
		})
	}
}
