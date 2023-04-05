package camera

import "github.com/Shamanskiy/go-ray-tracer/src/core"

type CameraSettings struct {
	VerticalFOV      core.Real
	AspectRatio      core.Real
	ImagePixelHeight int

	LookFrom core.Vec3
	LookAt   core.Vec3

	Antialiasing int

	ProgressChan chan<- int

	//float lensRadius{0.0};
}

func DefaultCameraSettings() CameraSettings {
	return CameraSettings{
		VerticalFOV:      90,
		AspectRatio:      2.,
		ImagePixelHeight: 360,
		LookFrom:         core.NewVec3(0., 0., 0.),
		LookAt:           core.NewVec3(0., 0., -1.),
		Antialiasing:     4,
	}
}
