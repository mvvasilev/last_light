package engine

import (
	"math"
	"math/rand"
)

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

func (wp *Positioned) SetPosition(pos Position) {
	wp.pos = pos
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

func (p Position) DistanceSquared(pos Position) int {
	return (pos.x-p.x)*(pos.x-p.x) + (pos.y-p.y)*(pos.y-p.y)
}

func (p Position) Distance(pos Position) int {
	return int(math.Floor(math.Sqrt(float64(p.DistanceSquared(pos)))))
}

func (p Position) Equals(other Position) bool {
	return p.x == other.x && p.y == other.y
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

func (s Size) Contains(x, y int) bool {
	return 0 <= x && x < s.width && 0 <= y && y < s.height
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
	if min == max {
		return min
	}

	return min + rand.Intn(max-min)
}

func MapSlice[S ~[]E, E any, R any](slice S, mappingFunc func(e E) R) []R {
	newSlice := make([]R, 0, len(slice))

	for _, el := range slice {
		newSlice = append(newSlice, mappingFunc(el))
	}

	return newSlice
}

func AbsInt(val int) int {
	switch {
	case val < 0:
		return -val
	case val == 0:
		return 0
	}
	return val
}
