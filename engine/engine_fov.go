package engine

import "math"

/* Stolen and modified from the go-fov package */

// Compute takes a GridMap implementation along with the x and y coordinates representing a player's current
// position and will internally update the visibile set of tiles within the provided radius `r`
func ComputeFOV[T any](transform func(x, y int) T, isInBounds, isOpaque func(x, y int) bool, px, py, radius int) (visibilityMap map[Position]T) {
	visibilityMap = make(map[Position]T)

	visibilityMap[PositionAt(px, py)] = transform(px, py)

	for i := 1; i <= 8; i++ {
		fov(visibilityMap, transform, isInBounds, isOpaque, px, py, 1, 0, 1, i, radius)
	}

	return visibilityMap
}

// fov does the actual work of detecting the visible tiles based on the recursive shadowcasting algorithm
// annotations provided inline below for (hopefully) easier learning
func fov[T any](visibilityMap map[Position]T, transform func(x, y int) T, isInBounds, isOpaque func(x, y int) bool, px, py, dist int, lowSlope, highSlope float64, oct, rad int) {
	// If the current distance is greater than the radius provided, then this is the end of the iteration
	if dist > rad {
		return
	}

	// Convert our slope into integers that will represent the "height" from the player position
	// "height" will alternately apply to x OR y coordinates as we move around the octants
	low := math.Floor(lowSlope*float64(dist) + 0.5)
	high := math.Floor(highSlope*float64(dist) + 0.5)

	// inGap refers to whether we are currently scanning non-blocked tiles consecutively
	// inGap = true means that the previous tile examined was empty
	inGap := false

	for height := low; height <= high; height++ {
		// Given the player coords and a distance, height and octant, determine which tile is being visited
		mapx, mapy := distHeightXY(px, py, dist, int(height), oct)
		if isInBounds(mapx, mapy) && distTo(px, py, mapx, mapy) < rad {
			// As long as a tile is within the bounds of the map, if we visit it at all, it is considered visible
			// That's the efficiency of shadowcasting, you just dont visit tiles that aren't visible
			visibilityMap[PositionAt(mapx, mapy)] = transform(mapx, mapy)
		}

		if isInBounds(mapx, mapy) && isOpaque(mapx, mapy) {
			if inGap {
				// An opaque tile was discovered, so begin a recursive call
				fov(visibilityMap, transform, isInBounds, isOpaque, px, py, dist+1, lowSlope, (height-0.5)/float64(dist), oct, rad)
			}
			// Any time a recursive call is made, adjust the minimum slope for all future calls within this octant
			lowSlope = (height + 0.5) / float64(dist)
			inGap = false
		} else {
			inGap = true
			// We've reached the end of the scan and, since the last tile in the scan was empty, begin
			// another on the next depth up
			if height == high {
				fov(visibilityMap, transform, isInBounds, isOpaque, px, py, dist+1, lowSlope, highSlope, oct, rad)
			}
		}
	}
}

// distHeightXY performs some bitwise and operations to handle the transposition of the depth and height values
// since the concept of "depth" and "height" is relative to whichever octant is currently being scanned
func distHeightXY(px, py, d, h, oct int) (int, int) {
	if oct&0x1 > 0 {
		d = -d
	}
	if oct&0x2 > 0 {
		h = -h
	}
	if oct&0x4 > 0 {
		return px + h, py + d
	}
	return px + d, py + h
}

// distTo is simply a helper function to determine the distance between two points, for checking visibility of a tile
// within a provided radius
func distTo(x1, y1, x2, y2 int) int {
	vx := math.Pow(float64(x1-x2), 2)
	vy := math.Pow(float64(y1-y2), 2)
	return int(math.Sqrt(vx + vy))
}
