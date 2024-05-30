package engine

func CastRay(pos1, pos2 Position) (points []Position) {
	x1, y1 := pos1.XY()
	x2, y2 := pos2.XY()

	isSteep := AbsInt(y2-y1) > AbsInt(x2-x1)

	if isSteep {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
	}

	reversed := false
	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
		reversed = true
	}

	deltaX := x2 - x1
	deltaY := AbsInt(y2 - y1)
	err := deltaX / 2
	y := y1
	var ystep int

	if y1 < y2 {
		ystep = 1
	} else {
		ystep = -1
	}

	for x := x1; x < x2+1; x++ {
		if isSteep {
			points = append(points, Position{y, x})
		} else {
			points = append(points, Position{x, y})
		}
		err -= deltaY
		if err < 0 {
			y += ystep
			err += deltaX
		}
	}

	if reversed {
		for i, j := 0, len(points)-1; i < j; i, j = i+1, j-1 {
			points[i], points[j] = points[j], points[i]
		}
	}

	return
}
