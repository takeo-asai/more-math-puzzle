package q6

type Rectangle struct {
	a int
	b int
}

func NewRectangle(x, y int) *Rectangle {
	if x >= y {
		return &Rectangle{x, y}
	}
	return &Rectangle{y, x}
}

func (r *Rectangle) Cut() *Rectangle {
	if r.a == r.b {
		return nil
	}
	return NewRectangle(r.a-r.b, r.b)
}
