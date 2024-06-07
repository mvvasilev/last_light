package engine

import (
	"log"
	"testing"
)

func BenchmarkPathfinding(b *testing.B) {
	path := FindPath(
		PositionAt(0, 0),
		PositionAt(16, 16),
		20,
		func(x, y int) bool {
			if x > 6 && x <= 16 && y == 10 {
				return false
			}

			if x < 0 || y < 0 {
				return false
			}

			if x > 16 || y > 16 {
				return false
			}

			return true
		},
	)

	if path == nil {
		log.Fatalf("No path could be found")
	}
}
