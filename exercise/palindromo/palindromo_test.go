package palindromo

import (
	"testing"
)

func TestPalindromo(t *testing.T) {

	tests := []struct {
		value         string
		valueExpected bool
	}{
		{"aaa", true},
		{"bbb", true},
	}

	for _, test := range tests {
		if PalindromoString(test.value) != test.valueExpected {
			t.Logf("Check this %s - exptected: %s", test.value, test.valueExpected)
		}
	}
}
