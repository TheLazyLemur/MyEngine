package main

import (
	"myengine/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UI struct {
	*core.Entity
	hasAwoken bool

	Show bool
}

func (ui *UI) SetEntity(e *core.Entity) {
	ui.Entity = e
}

func (ui *UI) GetTag() string {
	return "UI"
}
func (ui *UI) HasAwoken() bool {
	return ui.hasAwoken
}
func (ui *UI) Awake() {
	ui.hasAwoken = true
}

func (ui *UI) Update() {
	if !ui.Show {
		return
	}

	rl.DrawText("You win", 0, 0, 64, rl.Red)
}

func (ui *UI) HandleCollision(_ *core.Entity) {}

func (ui *UI) Toggle() {
	ui.Show = !ui.Show
}
