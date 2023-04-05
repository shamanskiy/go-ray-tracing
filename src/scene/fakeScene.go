package scene

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
)

type FakeScene struct {
	RecordedRays []core.Ray
}

func NewFakeScene() *FakeScene {
	return &FakeScene{}
}

func (fs *FakeScene) TestRay(ray core.Ray) color.Color {
	fs.RecordedRays = append(fs.RecordedRays, ray)
	return color.Red
}
