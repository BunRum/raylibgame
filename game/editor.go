package game

import (
	"encoding/json"
	"log"
	"os"
	"sdlgame/misc"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var edit bool
var selected *Object
var ogMouseX float32
var ogMouseY float32

var edits []*Object

func EditLoop() {
	mouseX := float32(rl.GetMouseX())
	mouseY := float32(rl.GetMouseY())
	if rl.IsKeyPressed(rl.KeyB) {
		edit = !edit
	}

	if edit {
		if rl.IsKeyPressed(rl.KeyF) {
			selected = Newobj(Object{Anchored: true, Collide: true})
			// edits = append(edits, selected)
		}
		if selected != nil {
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				selected.X = ogMouseX
				selected.Y = ogMouseY
				distanceX := mouseX - selected.X
				distanceY := mouseY - selected.Y
				println(distanceX, distanceY)

				if distanceX < 0 {
					selected.X += distanceX
					selected.W = -distanceX
				} else {
					selected.W = distanceX
				}

				if distanceY < 0 {
					selected.Y += distanceY
					selected.H = -distanceY
				} else {
					selected.H = distanceY
				}

			} else {
				ogMouseX = mouseX
				ogMouseY = mouseY
				selected.X = ogMouseX
				selected.Y = ogMouseY
			}
			if rl.IsKeyDown(rl.KeyUp) {
				selected.H += 1
			}
			if rl.IsKeyDown(rl.KeyDown) {
				selected.H = -1
			}
			if rl.IsKeyDown(rl.KeyLeft) {
				selected.W += -1
			}
			if rl.IsKeyDown(rl.KeyRight) {
				selected.W += 1
			}
		}
	}

}

func writetoJson() {
	var files misc.Typefiles
	err := misc.Readjsonfile("levels/level.json", &files)
	if err != nil {
		log.Fatalf("Error reading JSON file: %s", err)
	}
	jsonString, _ := json.Marshal(edits)
	os.WriteFile("big_marhsall.json", jsonString, os.ModePerm)
}

func AltEdit() {
	mouseX := float32(rl.GetMouseX())
	mouseY := float32(rl.GetMouseY())
	MouseRect := rl.NewRectangle(mouseX, mouseY, 10, 10)

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		if selected == nil {
			collision := RectCheckCollisions(MouseRect)
			if collision.Collided {
				selected = collision.Object
				edits = append(edits, selected)
			}
		}
		if selected != nil {
			distanceX := (mouseX - selected.W/2) - selected.X
			distanceY := (mouseY - selected.H/2) - selected.Y
			selected.VX = distanceX
			selected.VY = -distanceY

			if rl.IsKeyDown(rl.KeyUp) {
				selected.H += 1
			}
			if rl.IsKeyDown(rl.KeyDown) {
				selected.H = -1
			}
			if rl.IsKeyDown(rl.KeyLeft) {
				selected.W += -1
			}
			if rl.IsKeyDown(rl.KeyRight) {
				selected.W += 1
			}

		}
	}
	if rl.IsMouseButtonUp(rl.MouseLeftButton) {
		selected = nil
		writetoJson()
	}
}
