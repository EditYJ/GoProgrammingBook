package tempconv

func CToF(c Celsius) Fahreheit {
	return Fahreheit(c*9/5 + 32)
}
func FToC(f Fahreheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func KToC(k Kelvin) Celsius {
	return Celsius(k - KelvinZero)
}
