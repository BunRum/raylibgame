package game

import (
	_ "math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func calcSides(obj *Object) {
	obj.BottomY = obj.Y + obj.H
	obj.TopY = obj.Y - obj.H
	obj.LeftSideX = obj.X
	obj.RightSideX = obj.X + obj.W
	obj.MiddleX = obj.X + obj.W/2
	obj.MiddleY = obj.Y + obj.H/2
}

func checkBoundaries(obj *Object) {
	calcSides(obj)
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	if obj.RightSideX >= screenWidth || obj.X <= 0 {
		obj.VX = 0
		obj.X = obj.LastX
	}
	if obj.BottomY >= screenHeight || obj.Y <= 0 {
		obj.VY = 0
		obj.Y = obj.LastY
	}
}

type CollisionResult struct {
	Collided      bool
	Object        *Object
	CollisionRect rl.Rectangle
}

func RectCheckCollisions(rect rl.Rectangle) CollisionResult {
	for i := 0; i < len(objects); i++ {
		object := objects[i]
		if rl.CheckCollisionRecs(rect, object.Rect) {
			collisionRect := rl.GetCollisionRec(rect, object.Rect)
			rl.DrawRectangle(int32(collisionRect.X), int32(collisionRect.Y), int32(collisionRect.Width), int32(collisionRect.Height), rl.Green)
			return CollisionResult{Collided: true, Object: object, CollisionRect: collisionRect}
		}
	}
	return CollisionResult{}
}

func checkCollisions(obj *Object) CollisionResult {
	for i := 0; i < len(objects); i++ {
		otherobj := objects[i]
		if obj != otherobj && otherobj.Collide {
			rec1 := rl.Rectangle{X: obj.X, Y: obj.Y, Width: obj.W, Height: obj.H}
			rec2 := rl.Rectangle{X: otherobj.X, Y: otherobj.Y, Width: otherobj.W, Height: otherobj.H}
			if rl.CheckCollisionRecs(rec1, rec2) {
				println("they have collided")
				collisionRect := rl.GetCollisionRec(rec1, rec2)
				rl.DrawRectangle(int32(collisionRect.X), int32(collisionRect.Y), int32(collisionRect.Width), int32(collisionRect.Height), rl.Green)
				return CollisionResult{Collided: true, Object: otherobj, CollisionRect: collisionRect}
			}
		}
	}
	calcSides(obj)
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	if obj.RightSideX >= screenWidth || obj.X <= 0 || obj.BottomY >= screenHeight || obj.Y <= 0 {
		return CollisionResult{Collided: true}
	}
	return CollisionResult{Collided: false}
}

const Gravity = -1

func PhysicsLoop() {
	for i := 0; i < len(objects); i++ {
		obj := objects[i]
		obj.LastY = obj.Y
		obj.LastX = obj.X

		obj.VX *= 0.9

		obj.X += obj.VX

		// collisionX := checkCollisions(obj)
		// if collisionX != nil {
		// 	obj.VX = 0
		// 	obj.X = obj.LastX
		// }
		// obj.Y -= obj.VY
		// collisionY := checkCollisions(obj)
		// if !(collisionY == nil) {
		// 	obj.VY = 0
		// 	obj.Y = obj.LastY
		// }
		checkBoundaries(obj)

		if !obj.Anchored {
			obj.VY += -Gravity
		}

		obj.vars["lastvy"] = obj.VY
	}
}

// signum function returns 1 if n is positive, -1 if n is negative, and 0 if n is zero
func signum(n float32) float32 {
	switch {
	case n > 0:
		return 1
	case n < 0:
		return -1
	default:
		return 0
	}
}
func Altphys() {
	for i := 0; i < len(objects); i++ {
		object := objects[i]

		object.LastX = object.X
		object.LastY = object.Y

		// objects have gravity
		object.VY += Gravity
		// they have friction
		object.VX *= 0.9

		// X-axis collision
		object.X += object.VX
		collisionX := checkCollisions(object)
		if collisionX.Collided {
			// if math.Abs(float64(object.VX)) < 100 {
			// 	object.VX += signum(object.VX) * object.BounceSpeed
			// }
			// object.VX = -object.VX
			object.VX = 0
			object.X = object.LastX
			if object.OnCollide != nil {
				object.OnCollide()
			}
		}
		// Y-axis collision
		object.Y -= object.VY

		collisionY := checkCollisions(object)
		if collisionY.Collided {
			// if math.Abs(float64(object.VY)) < 100 {
			// 	object.VY += signum(object.VY) * object.BounceSpeed
			// }
			// object.VY = -object.VY
			if object.VY < -11 {
				object.VY = -object.VY / 2
			} else {
				object.VY = 0
			}
			object.Y = object.LastY
			if object.OnCollide != nil {
				object.OnCollide()
			}
		}
		// ------------------------------- \\

		object.Rect = rl.NewRectangle(object.X, object.Y, object.W, object.H)
	}
}
