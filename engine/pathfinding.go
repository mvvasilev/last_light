package engine

import (
	"slices"
)

type pathNode struct {
	pos Position

	parent *pathNode

	g int // distance between current node and start node
	h int // heuristic - squared distance from current node to end node
	f int // total cost of this node
}

func FindPath(from Position, to Position, isPassable func(x, y int) bool) *Path {
	var openList = make([]*pathNode, 0)
	var closedList = make([]*pathNode, 0)

	openList = append(openList, &pathNode{
		pos:    from,
		parent: nil,
		g:      0,
		h:      0,
		f:      0,
	})

	var lastNode *pathNode

	for {

		if len(openList) == 0 {
			break
		}

		// find node in open list with lowest f value, remove it from open and move it to closed
		currentIndex := 0
		currentNode := openList[currentIndex]

		for i, node := range openList {
			if node.f < currentNode.f {
				currentNode = node
				currentIndex = i
			}
		}

		if currentNode.pos.Equals(to) {
			lastNode = currentNode
			break // We have reached the goal
		}

		openList = slices.Delete(openList, currentIndex, currentIndex+1)
		closedList = append(closedList, currentNode)

		// use adjacent nodes as children
		children := []*pathNode{
			{
				pos:    currentNode.pos.WithOffset(1, 0),
				parent: currentNode,
			},
			{
				pos:    currentNode.pos.WithOffset(-1, 0),
				parent: currentNode,
			},
			{
				pos:    currentNode.pos.WithOffset(0, 1),
				parent: currentNode,
			},
			{
				pos:    currentNode.pos.WithOffset(0, -1),
				parent: currentNode,
			},
		}

		for _, child := range children {
			// If the child is impassable, skip it
			if !isPassable(child.pos.XY()) {
				continue
			}

			// If child is already contained in closedList, skip it
			if slices.ContainsFunc(closedList, func(el *pathNode) bool { return el.pos.Equals(child.pos) }) {
				continue
			}

			child.g = currentNode.g + 1
			child.h = to.DistanceSquared(child.pos)
			child.f = child.g + child.h

			// If child is already contained in openList, and has lower g
			if slices.ContainsFunc(openList, func(el *pathNode) bool { return el.pos.Equals(child.pos) && child.g > el.g }) {
				continue
			}

			openList = append(openList, child)
		}

	}

	if lastNode == nil {
		return nil
	}

	node := lastNode
	path := make([]Position, 0)

	for {
		path = append(path, node.pos)

		if node.parent == nil {
			break
		}

		node = node.parent
	}

	slices.Reverse(path)

	return CreatePath(
		from, to,
		path,
	)
}
