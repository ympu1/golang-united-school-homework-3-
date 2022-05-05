package homework

import "testing"

func TestReverse(t *testing.T) {
	testSlice := []int64{1, 2, 5, 15}
	got := reverse(testSlice)
	want := []int64{15, 5, 2, 1}
	if !ckeckEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func ckeckEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
