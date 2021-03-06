package octal

import (
	"testing"
)

var testCases = []struct {
	input       string
	expectedNum int64
	expectErr   bool
}{
	{"1", 1, false},
	{"10", 8, false},
	{"1234567", 342391, false},
	{"carrot", 0, true},
	{"35682", 0, true},
}

func TestParseOctal(t *testing.T) {
	for _, test := range testCases {
		actualNum, actualErr := ParseOctal(test.input)
		if actualNum != test.expectedNum {
			t.Fatalf("ParseOctal(%s): expected[%d], actual [%d]",
				test.input, test.expectedNum, actualNum)
		}

		// if we expect an error and there isn't one
		if test.expectErr && actualErr == nil {
			t.Errorf("ParseOctal(%s): expected an error, but error is nil", test.input)
		}
		// if we don't expect an error and there is one
		if !test.expectErr && actualErr != nil {
			t.Errorf("ParseOctal(%s): expected no error, but error is: %s", test.input, actualErr)
		}
	}
}

func BenchmarkParseOctal(b *testing.B) {
	b.StopTimer()

	for _, test := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ParseOctal(test.input)
		}

		b.StopTimer()
	}
}
