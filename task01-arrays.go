package homework

func average(input [15]float32) (result float32) {
	var j float32
	for i := 0; i < len(input); i++ {
		j = j + input[i]
	}
	return j / float32(len(input))
}
