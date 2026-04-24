package physics

type TransmissionType uint8
const (
	Manual TransmissionType = iota
	Automatic
)

type Transmission struct {
	trans 			TransmissionType
	gearRatios		[]float32 // ratio for each gear
	finalDriveRatio	float32
	currentGear		int8
	shiftPoints		[]float32
}
