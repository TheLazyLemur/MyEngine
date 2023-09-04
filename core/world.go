package core

import (
	"errors"
)

type World struct {
	Entities          []*Entity
	entitesForRemoval []int
}

func (w *World) AddEntity(e *Entity) {
	e.World = w
	w.Entities = append(w.Entities, e)
}

func (w *World) GetEntityWithTag(tag string) *Entity {
	for _, ent := range w.Entities {
		if ent.Tag == tag {
			return ent
		}
	}

	return nil
}

func (w *World) GetEntitiesWithTag(tag string) []*Entity {
	var ents []*Entity
	for _, ent := range w.Entities {
		if ent.Tag == tag {
			ents = append(ents, ent)
		}
	}

	return ents
}

func (e *World) GetComponentTag(cs []Component, tag string) (Component, error) {
	for _, c := range cs {
		if c.GetTag() == tag {
			return c, nil
		}
	}

	return nil, errors.New("Component not found")
}

func (e *World) Awake() {
	for _, e := range e.Entities {
		for _, c := range e.Components {
			if !c.HasAwoken() {
				c.Awake()
			}
		}
	}
}

func (w *World) Update() {
	for _, entR := range w.entitesForRemoval {
		for i, ent := range w.Entities {
			if ent.Id == entR {
				w.Entities = append(w.Entities[:i], w.Entities[i+1:]...)
			}
		}
	}

	w.entitesForRemoval = []int{}
	for _, e := range w.Entities {
		for _, c := range e.Components {
			c.Update()
		}
	}
}

func (e *World) HandleCollisions() {
	for _, x := range e.Entities {
		for _, y := range e.Entities {
			if x.Id != y.Id {
				rec1 := Rectangle{
					X:      x.Transform.Position.X,
					Y:      x.Transform.Position.Y,
					Width:  50,
					Height: 50,
				}
				rec2 := Rectangle{
					X:      y.Transform.Position.X,
					Y:      y.Transform.Position.Y,
					Width:  50,
					Height: 50,
				}

				if CheckCollisionRecs(rec1, rec2) {
					x.HandleCollision(y)
				}
			}
		}
	}
}

func (e *World) Destory(id int) {
	e.entitesForRemoval = append(e.entitesForRemoval, id)
}
