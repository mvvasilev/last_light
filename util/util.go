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
