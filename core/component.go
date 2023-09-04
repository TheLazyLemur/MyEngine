package core

type Component interface {
	SetEntity(e *Entity)
	GetTag() string
	HasAwoken() bool
	Awake()
	Update()
	HandleCollision(*Entity)
}
