package ui

type Point struct {
	X byte
	Y byte
}

// Add adds two points together and returns the resulting point
func (p *Point) Add(other *Point) *Point {
	return &Point{
		p.X + other.X,
		p.Y + other.Y,
	}
}
