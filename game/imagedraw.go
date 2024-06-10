package game

import (
	"image"

	// "image/color"
	"image/color"
	"image/png"

	// "strings"

	// "io"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type imagestuff struct {
	image  image.Image
	w      int
	h      int
	colors map[int]map[int]*color.RGBA
	testcolor [][]*color.RGBA
}

// var loadedimages = make(map[string]image.Image)
var testa = make(map[string]*imagestuff)

func Test() {
	// loadedfiles := make(map[string]*os.File)
	if testa["guy"] == nil {
		pngfile, err := os.Open("game/guy.png")
		if err != nil {
			panic(err)
			// return
		}
		image, err1 := png.Decode(pngfile)
		if err1 != nil {
			panic(err1)
			// return
		}
		// loadedimages["guy.png"] = image
		testa["guy"] = &imagestuff{
			image: image,
			w:     image.Bounds().Size().X,
			h:     image.Bounds().Size().Y,
		}
		testa["guy"].colors = make(map[int]map[int]*color.RGBA)
	}

	// guy := testa["guy"]
	row := 0
	xvalue := 0

	for unit := 0; unit <= testa["guy"].w*testa["guy"].h; unit++ {

		// if xvalue > imagew/2 {
		// var r, g, b, a uint32

		if testa["guy"].colors[row] == nil {
			// println("gskibidi bop")
			testa["guy"].colors[row] = make(map[int]*color.RGBA)
		}

		var coloraa color.RGBA
		if testa["guy"].colors[row][xvalue] == nil {
			r, g, b, a := testa["guy"].image.At(xvalue, row).RGBA()
			coloraa = color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}
			testa["guy"].colors[row][xvalue] = &coloraa
			// println(testa["guy"].colors[row][xvalue].A)
		} else {
			coloraa = *testa["guy"].colors[row][xvalue]
		}

		rl.DrawPixel(int32(xvalue)+500, int32(row)+200, coloraa)

		xvalue++
		// xvalue += rand.Intn(2)

		if xvalue >= testa["guy"].w {
			row++
			xvalue = 0
		}

	}
}

var textures = make(map[string]*rl.Texture2D)

func RenderImage() {
	if textures["guy"] == nil {
		imagea := rl.LoadImage("game/guy.png")
		tex := rl.LoadTextureFromImage(imagea)
		textures["guy"] = &tex
		rl.UnloadImage(imagea)
	}

	rl.DrawTexture(*textures["guy"], 400, 400, rl.White)
}
