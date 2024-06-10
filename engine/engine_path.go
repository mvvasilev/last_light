package engine

type Path struct {
	from Position
	to   Position

	path       []Position
	currentPos int
}

func CreatePath(from, to Position, path []Position) *Path {
	return &Path{
		from:       from,
		to:         to,
		path:       path,
		currentPos: 0,
	}
}

func (p *Path) From() Position {
	return p.from
}

func (p *Path) To() Position {
	return p.to
}

func (p *Path) CurrentPosition() Position {
	return p.path[p.currentPos]
}

func (p *Path) Next() (current Position, hasNext bool) {
	if p.currentPos+1 >= len(p.path) {
		return p.CurrentPosition(), false
	}

	p.currentPos++

	return p.CurrentPosition(), true
}

func LinePath(from, to Position) *Path {
	points := make([]Position, 0)
	n := float64(from.Distance(to))

	for step := 0.0; step <= n; step += 1.0 {
		t := 0.0

		if n != 0 {
			t = step / n
		}

		points = append(points, LerpPositions(from, to, t))
	}

	return CreatePath(from, to, points)
}
