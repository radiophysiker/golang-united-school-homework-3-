package homework

func sortMapValues(input map[int]string) (result []string) {
	var in = map[int]string{2: "a", 0: "b", 1: "c"}
	keys := make([]int, 0, len(in))
	for k := range in {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	array := make([]string, 0, len(in))
	for _, k := range keys {
		array = append(array, in[k])
	}
	return array
}
