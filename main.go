package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"math/rand"

	"myengine/core"
)

func main() {
	world := &core.World{
		Entities: []*core.Entity{},
	}

	playerEnt := &core.Entity{
		Id:  rand.Intn(1000),
		Tag: "Player",
		Transform: &core.Transform2D{
			Position: core.Vector2{
				X: -10,
				Y: 0,
			},
		},
	}
	playerEnt.AddComponent(&PlayerComponent{})

	enemyEnt := &core.Entity{
		Id:  rand.Intn(1000),
		Tag: "Enemy",
		Transform: &core.Transform2D{
			Position: core.Vector2{
				X: 60,
				Y: 0,
			},
		},
	}
	enemyEnt.AddComponent(&EnemyComponent{})

	enemyEnt2 := &core.Entity{
		Id:  rand.Intn(1000),
		Tag: "Enemy",
		Transform: &core.Transform2D{
			Position: core.Vector2{
				X: 120,
				Y: 0,
			},
		},
	}
	enemyEnt2.AddComponent(&EnemyComponent{})

	enemyEnt3 := &core.Entity{
		Id:  rand.Intn(1000),
		Tag: "Enemy",
		Transform: &core.Transform2D{
			Position: core.Vector2{
				X: 180,
				Y: 0,
			},
		},
	}
	enemyEnt3.AddComponent(&EnemyComponent{})

	uiEnt := &core.Entity{
		Id:  rand.Intn(1000),
		Tag: "UI",
		Transform: &core.Transform2D{
			Position: core.Vector2{
				X: 500,
				Y: 500,
			},
		},
	}
	uiEnt.AddComponent(&UI{})

	world.AddEntity(playerEnt)
	world.AddEntity(enemyEnt)
	world.AddEntity(enemyEnt2)
	world.AddEntity(enemyEnt3)
	world.AddEntity(uiEnt)

	rl.InitWindow(800, 600, "raylib [core] example - 2d camera")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(0, 0)

		world.Awake()
		world.HandleCollisions()
		world.Update()

		rl.EndDrawing()
	}
}
