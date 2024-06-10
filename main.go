// go:build main
package main

import (
	// "encoding/json"

	game "sdlgame/game"
	"sdlgame/misc"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

var files misc.Typefiles

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "raylib project")
	rl.InitAudioDevice()
	game.LoadAll()

	defer rl.CloseAudioDevice()
	defer rl.CloseWindow()
	defer game.UnloadAll()

	rl.SetTargetFPS(60)

	// err := misc.Readjsonfile("levels/level.json", &files)
	// if err != nil {
	// 	log.Fatalf("Error reading JSON file: %s", err)
	// }
	// println("this the shit", files, len(files))
	// for i := 0; i < len(files); i++ {
	// 	file := files[i]
	// 	println(i)
	// 	game.Newobj(game.Object{X: file.Position.X, Y: file.Position.X, W: file.Size.W, H: file.Size.H, Name: file.Name})
	// }

	game.Init()

	for !rl.WindowShouldClose() {
		rl.DrawFPS(0, 0)
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// game.EditLoop()
		game.AltEdit()
		game.GameLoop()
		game.Altphys()
		game.Draw()
		game.Test()
		// game.RenderImage()

		rl.EndDrawing()
		// rl.WaitTime(0.1)
	}
}
