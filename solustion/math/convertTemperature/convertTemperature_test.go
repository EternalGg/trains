package convertTemperature

func ConvertTemperature(celsius float64) []float64 {
	qa := make([]float64, 2)
	qa[0] = celsius + 273.15
	qa[1] = (celsius * 1.8) + 32
	return qa
}
