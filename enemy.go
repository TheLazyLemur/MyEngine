package main

import (
	// "fmt"
	"myengine/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyComponent struct {
	*core.Entity
	hasAwoken bool

	player *core.Entity
}

func (c *EnemyComponent) HasAwoken() bool {
	return c.hasAwoken
}

func (c *EnemyComponent) Awake() {
	c.hasAwoken = true
	c.player = c.GetEntityWithTag("Player")
}

func (c *EnemyComponent) Update() {
	c.Transform.Position.X += 10 * rl.GetFrameTime()
	// fmt.Printf("[ENEMY] Player with tag %s Position is (%v, %v)\n", c.player.Tag, c.player.transform.Position.X, c.player.transform.Position.Y)
	rl.DrawRectangle(int32(c.Transform.Position.X), int32(c.Transform.Position.Y), 50, 50, rl.Red)
}

func (c *EnemyComponent) HandleCollision(e *core.Entity) {
	// fmt.Println("ENEMY collided with", e.Tag)
	if e.Tag == "Player" {
		c.World.Destory(c.Id)
	}
}

func (c *EnemyComponent) GetTag() string {
	return "Enemy"
}

func (c *EnemyComponent) SetEntity(e *core.Entity) {
	c.Entity = e
}
