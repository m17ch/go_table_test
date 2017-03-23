package slowtest

import "testing"

var slowTests = []struct {
	input    string // input
	expected bool   // expected result
}{
	{"just", true},
	{"playing", true},
	{"with", true},
	{"random", true},
	{"sleep", true},
	{"for", true},
	{"test", true},
	{"duration", true},
	{"variety", true},
}

func TestSlow(t *testing.T) {
	for _, tt := range slowTests {
		actual := Slow(tt.input)
		if actual != tt.expected {
			t.Errorf("Slow(%s): expected %t, actual %t", tt.input, tt.expected, actual)
		}
	}
}
