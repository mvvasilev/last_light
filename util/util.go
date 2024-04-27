package util

type Position struct {
	x int
	y int
}

func PositionAt(x int, y int) Position {
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

func (p Position) WithOffset(xOffset int, yOffset int) Position {
	p.x = p.x + xOffset
	p.y = p.y + yOffset
	return p
}

type Size struct {
	width  int
	height int
}

func SizeOf(width int, height int) Size {
	return Size{int(width), int(height)}
}

func SizeOfInt(width int, height int) Size {
	return Size{width, height}
}

func (s Size) Width() int {
	return s.width
}

func (s Size) Height() int {
	return s.height
}

func (s Size) WH() (int, int) {
	return s.width, s.height
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
