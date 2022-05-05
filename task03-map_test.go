package homework

import "testing"

func TestSortMapValues(t *testing.T) {
	input := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	got := sortMapValues(input)
	want := []string{"bb", "aa", "cc"}
	if !checkStringSLiceEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func checkStringSLiceEqual(a, b []string) bool {
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
