package main

import (
	"fmt"

	tmp "units/temperature"
)

func main() {
	celsius := tmp.ToCelsius(95)
	fahrenheit := tmp.ToFahrenheit(35)

	fmt.Println(35, "grados celsius corresponden a", fahrenheit, "grados fahrenheit")
	fmt.Println(95, "grados fahrenheit corresponden a", celsius, "grados celsius")
}
