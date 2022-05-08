package homework

func reverse(input []int64) (result []int64) {
	var inverse []int64
	for i := len(input) - 1; i >= 0; i-- {
		inverse = append(inverse, input[i])
	}
	return inverse
}
