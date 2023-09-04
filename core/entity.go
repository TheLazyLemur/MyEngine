package core

type Entity struct {
	Id         int
	Tag        string
	Components []Component
	World      *World

	Transform *Transform2D
}

func (e *Entity) AddComponent(c Component) {
	c.SetEntity(e)
	e.Components = append(e.Components, c)
}

func (e *Entity) GetEntityWithTag(tag string) *Entity {
	return e.World.GetEntityWithTag(tag)
}

func (e *Entity) GetEntitiesWithTag(tag string) []*Entity {
	return e.World.GetEntitiesWithTag(tag)
}

func (e *Entity) GetComponentTag(tag string) (Component, error) {
	return e.World.GetComponentTag(e.Components, tag)
}

func (e *Entity) HandleCollision(other *Entity) {
	for _, c := range e.Components {
		c.HandleCollision(other)
	}
}
