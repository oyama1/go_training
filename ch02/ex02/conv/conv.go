package conv

const(
	Temp = "temp"
	Length = "length"
	Weight = "weight"
)

// 摂氏の温度を華氏へ変換します
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// 摂氏の温度を絶対温度へ変換します
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

// 華氏の温度を摂氏へ変換します
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// メートル(m)をフィート(f)へ変換します
func MeterToFeet(m Meter) Feet {
	return Feet(m * 3.2808)
}

// フィート(f)をメートル(m)へ変換します
func FeetToMeter(f Feet) Meter {
	return Meter(f * 0.3048)
}

// キログラム(kg)をポンド(lb)へ変換します
func KgToLb(kg Kg) Lb {
	return Lb(kg * 2.2046)
}

// ポンド(lb)をキログラム(kg)へ変換します
func LbToKg(lb Lb) Kg {
	return Kg(lb * 0.4536)
}
