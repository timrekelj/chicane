package main

import (
	phys "chicane/physics"
	sys "chicane/systems"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	vehicle phys.Vehicle
}

func main() {
	rl.InitWindow(720, 360, "Chicane")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var game = Game{}
	game.vehicle = phys.NewVehicle()

	for !rl.WindowShouldClose() {
		var _ = rl.GetFrameTime()

		// Input system
		input := sys.ReadInput()
		game.vehicle.ApplyInput(input)

		// Render
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		sys.RenderDashboard(game.vehicle.Dashboard())

		rl.EndDrawing()
	}
}
