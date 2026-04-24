package systems

import (
	phys "chicane/physics"

	rl "github.com/gen2brain/raylib-go/raylib"

	"fmt"
)

func RenderDashboard(dash phys.DashboardData) {
	rl.DrawText(fmt.Sprintf("Speed: %.2f km/h", dash.SpeedKPH), 10, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("RPM: %.2f", dash.RPM), 10, 40, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Gear: %d", dash.Gear), 10, 70, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Throttle: %.2f", dash.Throttle), 10, 100, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Brake: %.2f", dash.Brake), 10, 130, 20, rl.White)
}
