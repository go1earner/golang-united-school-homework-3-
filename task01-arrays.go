package homework

func average(input [15]float32) (result float32) {
	var sum float32
	for _, i := range input {
		sum += i
	}
	result = sum / 15.0
	return
}
