package util

import "math/rand"

type Positioned struct {
	pos Position
}

func WithPosition(pos Position) Positioned {
	return Positioned{
		pos: pos,
	}
}

func (wp *Positioned) Position() Position {
	return wp.pos
}

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

type Sized struct {
	size Size
}

func WithSize(size Size) Sized {
	return Sized{
		size: size,
	}
}

// Checks if the provided coordinates fit within the sized struct, [0, N)
func (ws *Sized) FitsWithin(x, y int) bool {
	return 0 <= x && x < ws.size.width && 0 <= y && y < ws.size.height
}

func (ws *Sized) Size() Size {
	return ws.size
}

type Size struct {
	width  int
	height int
}

func SizeOf(width int, height int) Size {
	return Size{int(width), int(height)}
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

func (s Size) Area() int {
	return s.width * s.height
}

func (s Size) AsArrayIndex(x, y int) int {
	return y*s.width + x
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

func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}

type Room struct {
	Positioned
	Sized
}
