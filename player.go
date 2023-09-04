package main

import (
	"fmt"

	"myengine/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerComponent struct {
	*core.Entity
	hasAwoken bool

	enemy *core.Entity
	UI    *UI
	win   bool
}

func (c *PlayerComponent) HasAwoken() bool {
	uiEnt := c.GetEntityWithTag("UI")
	uiComp, err := uiEnt.GetComponentTag("UI")
	if err != nil {
		panic(err)
	}

	c.UI = uiComp.(*UI)
	return c.hasAwoken
}

func (c *PlayerComponent) HandleCollision(other *core.Entity) {
	fmt.Println("PLAYER collided with", other.Tag, other.Id)
}

func (c *PlayerComponent) Awake() {
	c.hasAwoken = true

	c.enemy = c.GetEntityWithTag("Enemy")
	if c.enemy == nil {
		panic("Enemy not found")
	}
}

func (c *PlayerComponent) Update() {
	enemies := c.GetEntitiesWithTag("Enemy")
	if len(enemies) <= 0 && c.win == false {
		c.UI.Toggle()
		c.win = true
	}

	if rl.IsKeyDown(rl.KeyA) {
		c.Transform.Position.X -= 100 * rl.GetFrameTime()
	}

	if rl.IsKeyDown(rl.KeyD) {
		c.Transform.Position.X += 100 * rl.GetFrameTime()
	}

	if rl.IsKeyDown(rl.KeyW) {
		c.Transform.Position.Y -= 100 * rl.GetFrameTime()
	}

	if rl.IsKeyDown(rl.KeyS) {
		c.Transform.Position.Y += 100 * rl.GetFrameTime()
	}

	rl.DrawRectangle(int32(c.Transform.Position.X), int32(c.Transform.Position.Y), 50, 50, rl.Blue)
}

func (c *PlayerComponent) GetTag() string {
	return "Player"
}

func (c *PlayerComponent) SetEntity(e *core.Entity) {
	c.Entity = e
}
