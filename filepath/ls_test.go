package filepath

import (
	"testing"
)

func TestLsByGlob(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "Glob"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LsByGlob()
		})
	}
}

func TestLsByDir(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "ioutil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LsByDir()
		})
	}
}
