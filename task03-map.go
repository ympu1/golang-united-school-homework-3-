package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var keys []int
	for key := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		result = append(result, input[key])
	}
	return
}
