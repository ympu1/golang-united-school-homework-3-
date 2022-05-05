package homework

import "testing"

func TestAverage(t *testing.T) {
	testSlice := [15]float32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	got := average(testSlice)
	var want float32 = 1.0
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
