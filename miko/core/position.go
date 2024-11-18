package core

type Position struct {
	X float64
	Y float64
}

type Box struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

func PositionZero() *Position {
	return &Position{
		X: 0,
		Y: 0,
	}
}
func Position64() *Position {
	return &Position{
		X: 64,
		Y: 0,
	}
}
func Position128() *Position {
	return &Position{
		X: 128,
		Y: 0,
	}
}
func Position192() *Position {
	return &Position{
		X: 192,
		Y: 0,
	}
}
