package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Player *Object

var Grabbed bool

func Init() {
	Player = Newobj(Object{W: 20, H: 20, Color: rl.Gold}.Center())
	Player.vars["tag"] = "player"
}

func GameLoop() {

	if rl.IsKeyPressed(rl.KeyG) {
		Newobj(Object{W: 50, H: 50, VX: 1, VY: 1, BounceSpeed: 2, Kind: "Circle"}.Center())
	}

	if rl.IsKeyDown(rl.KeyA) {
		Player.VX = -10.0
	}
	if rl.IsKeyDown(rl.KeyD) {
		Player.VX = 10.0
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		Player.VY = 10
		rl.PlaySound(sounds["jump"])
	}

	rl.DrawText(fmt.Sprintf("VY: %f", Player.VY), 400, 20, 20, rl.Green)
	// rl.DrawText(fmt.Sprintf("abs VY: %f", math.Abs(float64(Player.VY))), 400, 40, 20, rl.Green)
	rl.DrawText(fmt.Sprintf("VX: %f", Player.VX), 400, 60, 20, rl.Green)
}
