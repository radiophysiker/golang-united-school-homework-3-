package homework

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	array := make([]string, 0, len(input))
	for _, k := range keys {
		array = append(array, input[k])
	}
	return array
}
