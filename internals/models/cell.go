package models

import (
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	rl.Vector2
	Radius float32
	Color  rl.Color
}

func random(maximum float32) float32 { return rand.Float32() * maximum }
func NewCell() *Cell {
	cell := &Cell{
		Radius: 20,
	}
	randFrom := func(min, max uint8) uint8 {
		return uint8(rand.Intn(int(max-min+1)) + int(min))
	}
	cell.X = random(float32(rl.GetScreenWidth()))
	cell.Y = random(float32(rl.GetScreenHeight()))
	cell.Color = rl.NewColor(randFrom(100, 255), 0, randFrom(100, 255), 100)
	return cell
}

func (c *Cell) Draw() {
	rl.DrawCircleV(c.Vector2, c.Radius, c.Color)
}

func (c *Cell) Clicked(v rl.Vector2) bool {
	fmt.Println("Checking collision")
	fmt.Println(c.Vector2)
	return rl.CheckCollisionPointCircle(v, c.Vector2, c.Radius)
}

func (c *Cell) Mitosis() []Cell {
	cell := NewCell()
	cell.Vector2 = c.Vector2
	cell.Color = c.Color
	cell.Radius = c.Radius * .8
	c.X += rand.Float32() * cell.Radius * 2
	cellB := NewCell()
	cellB.Vector2 = c.Vector2
	cellB.Color = c.Color
	cellB.Radius = c.Radius * .8
	c.X += rand.Float32() * cell.Radius * 2

	return []Cell{*cell, *cellB}
}

func random2D(v *rl.Vector2) {
	angle := float64(random(2 * math.Pi))
	v.X = float32(math.Cos(angle))
	v.Y = float32(math.Sin(angle))
}

func (c *Cell) Move() {
	vel := rl.Vector2{}
	random2D(&vel)
	c.Vector2 = rl.Vector2Add(c.Vector2, vel)
}
