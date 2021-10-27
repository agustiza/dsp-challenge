package budget

type currency int64

func Currency(f float64) currency {
	return currency((f * 100) + 0.5)
}

func (m currency) float() float64 {
	x := float64(m)
	x = x / 100
	return x
}

func (m currency) multiply(f float64) currency {
	x := (float64(m) * f) + 0.5
	return currency(x)
}
