package util

type Position struct {
	x int
	y int
}

func PositionAt(x uint16, y uint16) Position {
	return Position{int(x), int(y)}
}

func (p Position) X() int {
	return p.x
}

func (p Position) Y() int {
	return p.y
}

func (p Position) XY() (int, int) {
	return p.x, p.y
}

func (p Position) XYUint16() (uint16, uint16) {
	return uint16(p.x), uint16(p.y)
}

type Size struct {
	width  int
	height int
}

func SizeOf(width uint16, height uint16) Size {
	return Size{int(width), int(height)}
}

func (s Size) Width() int {
	return s.width
}

func (s Size) Height() int {
	return s.height
}

func (s Size) WHUint16() (uint16, uint16) {
	return uint16(s.width), uint16(s.height)
}

func LimitIncrement(i int, limit int) int {
	if (i + 1) > limit {
		return i
	}

	return i + 1
}

func LimitDecrement(i int, limit int) int {
	if (i - 1) < limit {
		return i
	}

	return i - 1
}
