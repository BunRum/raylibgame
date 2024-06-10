package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var time float32

func Draw() {
	for i := 0; i < len(objects); i++ {
		obj := objects[i]
		calcSides(obj)

		// rl.Translatef(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2), 0)
		// rl.Rotatef(time*50, 0, 0, -1)
		// rl.DrawRectangleLines(int32(obj.X), int32(obj.Y), int32(obj.W), int32(obj.H), rl.White)
		// color := if obj.vars["color"]
		// color.RGBA{}
		var col color.RGBA
		if obj.Color != (color.RGBA{}) {
			col = obj.Color
		} else {
			col = rl.Green
		}
		rl.DrawRectanglePro(obj.Rect, rl.Vector2{X: 0, Y: 0}, 0, col)
		// rl.DrawRectangle(int32(obj.X), int32(obj.Y), int32(obj.W), int32(obj.H), rl.White)
		// rl.DrawCircle(int32(obj.LeftSideX), int32(obj.Y), 5, rl.Green)
		// rl.DrawCircle(int32(obj.LeftSideX), int32(obj.BottomY), 5, rl.Green)
		// rl.DrawCircle(int32(obj.RightSideX), int32(obj.Y), 5, rl.Green)
		// rl.DrawCircle(int32(obj.RightSideX), int32(obj.BottomY), 5, rl.Green)
		// rl.DrawCircle(int32(obj.MiddleX), int32(obj.MiddleY), 5, rl.Green)
		time += 0.5
	}
}
