package quote

import "testing"

func TestSunAlso(t *testing.T) {
	i := interface{}(SunAlso)
	_, ok := i.(string)
	if !ok {
		t.Errorf("Expected SunAlso to be a string")
	}
}
