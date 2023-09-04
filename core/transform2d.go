package core

type Transform2D struct {
	Position Vector2
	Rotation float64
	Scale    Vector2
}

type Vector2 struct {
	X float32
	Y float32
}
