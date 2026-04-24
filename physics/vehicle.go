package physics

type VehicleInput struct {
	Throttle	float32
	Brake		float32
	Steering	float32
	GearChange	int8 // -1 downshift, 1 upshift, 0 no change
	Clutch		bool
	Handbrake	bool
}

type Vehicle struct {
	engine			Engine
	trans			Transmission
	brakeForce		float32
	mass			float32 // kg
	velocity		float32 // m/s
	tires			[4]Tire
	wheelsLocked	bool
	driveWheels		[]int // [2, 3] for rear wheel drive
}

type DashboardData struct {
	SpeedKPH	float32
	RPM			float32
	Gear		int8
	Throttle	float32
	Brake		float32
}

func (v *Vehicle) ApplyInput(input VehicleInput) {
	v.engine.throttle = input.Throttle
	v.brakeForce = input.Brake

	newGear := int(v.trans.currentGear) + int(input.GearChange)
	if newGear >= -1 && newGear < len(v.trans.gearRatios) {
		v.trans.currentGear = int8(newGear)
	}

	v.wheelsLocked = input.Handbrake
}

func (v *Vehicle) Update(dt float32) {
	// all car physics will happen here
}

func (v *Vehicle) Dashboard() DashboardData {
	return DashboardData{
		SpeedKPH: v.velocity * 3.6,
		RPM: v.engine.currentRPM,
		Gear: v.trans.currentGear,
		Throttle: v.engine.throttle,
		Brake: v.brakeForce,
	}
}

func NewVehicle() Vehicle {
	return Vehicle{
		engine: Engine{
			torqueCurve: []float32{150, 180, 210, 220, 215, 200, 175},
			powerCurve: []float32{150, 180, 210, 220, 215, 200, 175},
			maxRPM: 7500,
			redlineRPM: 7000,
			idleRPM: 800,

			currentRPM: 0,
			throttle: 0,
			inertia: 0,
		},
		trans: Transmission{
			trans: Manual,
			gearRatios: []float32{0, 3.14, 2.05, 1.48, 1.13, 0.89, 0.72},
			finalDriveRatio: 4.1,

			currentGear: 0,
			shiftPoints: []float32{},
		},
		tires: [4]Tire{
			{ radius: 0.31, gripCoeff: 0.9, normalForce: 0, slipRatio: 0, },
			{ radius: 0.31, gripCoeff: 0.9, normalForce: 0, slipRatio: 0, },
			{ radius: 0.31, gripCoeff: 0.9, normalForce: 0, slipRatio: 0, },
			{ radius: 0.31, gripCoeff: 0.9, normalForce: 0, slipRatio: 0, },
		},
		mass: 1100,
		driveWheels: []int{2, 3},

		brakeForce: 0,
		velocity: 0,
		wheelsLocked: false,
	}
}
