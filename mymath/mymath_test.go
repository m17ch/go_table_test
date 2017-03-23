package mymath

import "testing"

var halvesTests = []struct {
	input    float64 // input
	expected float64 // expected result
}{
	{0, 0},
	{2, 1},
	{4, 2},
	{10, 5},
	{-5, -2.5},
	{-12, -6},
	{100, 50},
}

func TestHalves(t *testing.T) {
	for _, tt := range halvesTests {
		actual := Halves(tt.input)
		if actual != tt.expected {
			t.Errorf("Halves(%f): expected %f, actual %f", tt.input, tt.expected, actual)
		}
	}
}

var isEvenTests = []struct {
	input    int  // input
	expected bool // expected result
}{
	{-1491, false},
	{-5, false},
	{0, true},
	{2, true},
	{4, true},
	{5, false},
	{7, false},
	{20, true},
	{12456, true},
	{20853, false},
}

func TestIsEven(t *testing.T) {
	for _, tt := range isEvenTests {
		actual := IsEven(tt.input)
		if actual != tt.expected {
			t.Errorf("IsEven(%d): expected %t, actual %t", tt.input, tt.expected, actual)
		}
	}
}

var additionTests = []struct {
	input1   int // input 1
	input2   int // input 2
	expected int // expected result
}{
	{0, 1, 1},
	{-4, 3, -1},
	{1, 1, 2},
	{100, -100, 0},
	{4, 3, 7},
	{2000, 20, 2020},
	{0, 0, 0},
	{1, 1, 2},
	{5, 6, 11},
	{100, 100, 200},
	{-10000, 10000, 0},
}

func TestAddition(t *testing.T) {
	for _, tt := range additionTests {
		actual := Addition(tt.input1, tt.input2)
		if actual != tt.expected {
			t.Errorf("Addition(%d,%d): expected %d, actual %d", tt.input1, tt.input2, tt.expected, actual)
		}
	}
}
