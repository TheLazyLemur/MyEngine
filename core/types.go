package core

type Rectangle struct {
	X, Y, Width, Height float32
}

func CheckCollisionRecs(rec1, rec2 Rectangle) bool {
	right1 := rec1.X + rec1.Width
	bottom1 := rec1.Y + rec1.Height
	right2 := rec2.X + rec2.Width
	bottom2 := rec2.Y + rec2.Height

	if rec1.X < right2 && right1 > rec2.X &&
		rec1.Y < bottom2 && bottom1 > rec2.Y {
		return true
	}

	return false
}
