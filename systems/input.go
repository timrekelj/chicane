package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	phys "chicane/physics"
)

const (
	KeyThrottle		= rl.KeyW
	KeyBrake		= rl.KeyS
	KeyUpshift		= rl.KeyE
	KeyDownshift	= rl.KeyQ

	GamepadThrottle 	= rl.GamepadAxisRightTrigger
	GamepadBrake		= rl.GamepadAxisLeftTrigger
	GamepadUpshift		= rl.GamepadButtonRightFaceLeft
	GamepadDownshift	= rl.GamepadButtonRightFaceDown
)

func ReadInput() phys.VehicleInput {
	input := phys.VehicleInput{}

	if rl.IsKeyDown(KeyThrottle) {
		input.Throttle = 1
	} else {
		input.Throttle = 0
	}

	if rl.IsKeyDown(KeyBrake) {
		input.Brake = 1
	} else {
		input.Brake = 0
	}

	if rl.IsKeyPressed(KeyUpshift) {
		input.GearChange = 1
	} else if rl.IsKeyPressed(KeyDownshift) {
		input.GearChange = -1
	} else {
		input.GearChange = 0
	}

	if rl.IsGamepadAvailable(0) {
		// Default axis range is from -1 to 1

		if t := (rl.GetGamepadAxisMovement(0, GamepadThrottle) + 1) / 2; t > 0.05 {
			input.Throttle = t
		}
		if b := (rl.GetGamepadAxisMovement(0, GamepadBrake) + 1) / 2; b > 0.05 {
			input.Brake = b
		}

		if (input.GearChange == 0) {
			if rl.IsGamepadButtonPressed(0, GamepadUpshift) {
				input.GearChange = 1
			} else if rl.IsGamepadButtonPressed(0, GamepadDownshift) {
				input.GearChange = -1
			}
		}
	}

	return input
}
