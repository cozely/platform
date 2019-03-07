// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package window

import "math"

// Coord represents the coordinates of a pixel on the window.
type Coord struct {
	X, Y int32
}

// XY returns an integer vector corresponding to the first two coordinates of
// v.
func XY(x, y int32) Coord {
	return Coord{x, y}
}

// Round returns an integer vector corresponding to the first two
// coordinates of v.
func Round(x, y float64) Coord {
	return Coord{
		int32(math.Round(x)),
		int32(math.Round(y)),
	}
}

// Plus returns the sum with another vector.
func (a Coord) Plus(b Coord) Coord {
	return Coord{a.X + b.X, a.Y + b.Y}
}

// Minus returns the difference with another vector.
func (a Coord) Minus(b Coord) Coord {
	return Coord{a.X - b.X, a.Y - b.Y}
}

// Opposite returns the opposite of the vector.
func (a Coord) Opposite() Coord {
	return Coord{-a.X, -a.Y}
}

// Times returns the component-wise product with another vector.
func (a Coord) Times(b Coord) Coord {
	return Coord{a.X * b.X, a.Y * b.Y}
}

// Slash returns the integer quotients of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (a Coord) Slash(b Coord) Coord {
	return Coord{a.X / b.X, a.Y / b.Y}
}

// Mod returns the remainder (modulus) of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (a Coord) Mod(b Coord) Coord {
	return Coord{a.X % b.X, a.Y % b.Y}
}

// FlipX returns the vector with the sign of X flipped.
func (a Coord) FlipX() Coord {
	return Coord{-a.X, a.Y}
}

// FlipY returns the vector with the sign of Y flipped.
func (a Coord) FlipY() Coord {
	return Coord{a.X, -a.Y}
}

// ProjX returns the vector projected on the X axis (i.e. with Y nulled).
func (a Coord) ProjX() Coord {
	return Coord{a.X, 0}
}

// ProjY returns the vector projected on the Y axis (i.e. with X nulled).
func (a Coord) ProjY() Coord {
	return Coord{0, a.Y}
}

// YX returns the vector with coordinates X and Y swapped.
func (a Coord) YX() Coord {
	return Coord{a.Y, a.X}
}

// Perp returns the vector rotated by 90 in counter-clockwise direction.
func (a Coord) Perp() Coord {
	return Coord{-a.Y, a.X}
}

// Null returns true if both coordinates are null.
func (a Coord) Null() bool {
	return a.X == 0 && a.Y == 0
}
