package examplepando

import "testing"

func TestSumasEsto(t *testing.T) {
	esto := SumaEsto(1, 2)
	if esto != 3 {
		t.Fatal("error")
	}
}
