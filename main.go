package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type S_Vehicle struct {
	name		string
	color 		rl.Color
	size 		rl.Vector2
	position	rl.Vector2
	rotation	float32
	velocity	rl.Vector2
	tireColor	rl.Color
	tireOffset	[4]rl.Vector2
}

func (v *S_Vehicle) GetPosition() rl.Vector2 {
	return v.position
}

func (v *S_Vehicle) Update(dt float32) {
	engineForce := 500.0

	v.position.X += v.velocity.X * float32(engineForce) * dt
}

func (v *S_Vehicle) Draw() {
	rl.DrawRectangle(
		int32(v.position.X), int32(v.position.Y),
		int32(v.size.X), int32(v.size.Y),
		v.color,
	)

	velocityText := fmt.Sprintf("velocity: %.1f", rl.Vector2Length(v.velocity))
	rl.DrawText(velocityText, 10, 10, 20, rl.White)
}

func (v *S_Vehicle) accelerate() {
	v.velocity.X = v.velocity.X + 10;
}

func (v *S_Vehicle) Input() {
	if rl.IsKeyDown(rl.KeyW) {
		v.accelerate()
	}
}

type Game struct {
	vehicles []S_Vehicle
}

func main() {
	rl.InitWindow(720, 360, "Chicane")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var game = Game{}
	game.vehicles = append(game.vehicles, S_Vehicle{
		name: "RX7",
		color: rl.Blue,
		size: rl.Vector2{X: 40.00, Y: 100.00},
		position: rl.Vector2{X: 100, Y: 100},
		rotation: 0,
		velocity: rl.Vector2{X: 0, Y: 0},
		tireColor: rl.Black,
		tireOffset: [4]rl.Vector2{
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
		},
	})

	for !rl.WindowShouldClose() {
		var dt = rl.GetFrameTime()

		game.handleInput()
		game.updatePhysics(dt)
		game.render()
	}
}

func (g *Game) handleInput() {
	for i := range g.vehicles {
		g.vehicles[i].Input()
	}
}

func (g *Game) updatePhysics(dt float32) {
	for i := range g.vehicles {
		g.vehicles[i].Update(dt)
	}
}

func (g *Game) render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	for _, v := range g.vehicles {
		v.Draw()
	}

	rl.EndDrawing()
}
