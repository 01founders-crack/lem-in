package tests

import (
	"lem-in/helpers"
	"testing"
)

func TestContainsStartCharCheck(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"#hello", true},
		{"Lhello", true},
		{"hello", false},
		{"!hello", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := helpers.ContainsStartCharCheck(test.input)
			if actual != test.expected {
				t.Errorf("ContainsStartCharCheck(%q) = %v; want %v", test.input, actual, test.expected)
			}
		})
	}
}

func TestContainsSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"hello world", true},
		{"hello", false},
		{"hello ", true},
		{"", false},
		{" ", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := helpers.ContainsSpace(test.input)
			if actual != test.expected {
				t.Errorf("ContainsSpace(%q) = %v; want %v", test.input, actual, test.expected)
			}
		})
	}
}
