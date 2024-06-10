package game

import rl "github.com/gen2brain/raylib-go/raylib"

var soundfiles = map[string]string{
	"jump": "assets/jump.wav",
}

var sounds = make(map[string]rl.Sound)

func LoadAll() {
	for i, v := range soundfiles {
		println(i, v)
		// sounds[i] = rl.LoadSound(v)
		sounds[i] = rl.LoadSound(v)
		println(sounds)
	}
}

func UnloadAll() {
	for _, v := range sounds {
		rl.UnloadSound(v)
	}
}
