package convertTemperature

func convertTemperature(celsius float64) []float64 {
	result := make([]float64, 2)
	result[0] = celsius + 273.15
	result[1] = (celsius * 1.8) + 32
	return result
}
