package physics

type Engine struct {
	torqueCurve []float32 // torque at different RPM points
	powerCurve	[]float32 // power at different RPM points
	currentRPM	float32
	maxRPM		float32
	redlineRPM	float32
	idleRPM		float32
	throttle	float32 // 0.00 - 1.00
	inertia		float32
}
