package temperature

// ToCelsius ...
func ToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) * 5.0 / 9.0
}

// ToFahrenheit ...
func ToFahrenheit(celsius float64) float64 {
	return celsius*9.0/5.0 + 32.0
}
