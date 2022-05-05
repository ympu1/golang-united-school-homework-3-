package homework

import "testing"

func TestAverage(t *testing.T) {
	testSlice := []float32{1, 2, 3, 4, 5, 6}
	got := average(testSlice)
	var want float32 = 3.5
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
