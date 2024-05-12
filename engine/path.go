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
