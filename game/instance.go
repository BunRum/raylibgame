package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Object struct {
	Id          int
	Name string
	X           float32
	VX          float32
	LastX       float32
	Y           float32
	VY          float32
	LastY       float32
	W           float32
	H           float32
	Kind        string
	Anchored    bool
	Collide     bool
	BottomY     float32 `json:"-"`
	TopY        float32 `json:"-"`
	LeftSideX   float32 `json:"-"`
	RightSideX  float32 `json:"-"`
	MiddleX     float32 `json:"-"`
	MiddleY     float32 `json:"-"` 
	vars        map[string]any 
	BounceSpeed float32
	Rect        rl.Rectangle `json:"-"`
	OnCollide   Collider `json:"-"`
	Color       color.RGBA
}

type Collider func()

var objects []*Object

func Newobj(obj Object) *Object {
	obj.Id = len(objects)
	if obj.vars == nil {
		obj.vars = make(map[string]any)
	}
	if !obj.Collide {
		obj.Collide = true
	}

	if obj.BounceSpeed == 0 {
		obj.BounceSpeed = 1
	}
	
	obj.Rect = rl.NewRectangle(obj.X, obj.Y, obj.W, obj.H)

	objects = append(objects, &obj)
	return &obj
}

func (obj Object) Center() Object {
	obj.X = (float32(rl.GetScreenWidth()) / 2) - obj.W/2
	obj.Y = (float32(rl.GetScreenHeight()) / 2) - obj.H/2
	return obj
}
