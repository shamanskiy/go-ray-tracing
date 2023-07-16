package scene

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
)

type FakeScene struct {
	RecordedRays  []core.Ray
	ColorToReturn color.Color
}

func NewFakeScene(colorToReturn color.Color) *FakeScene {
	return &FakeScene{ColorToReturn: colorToReturn}
}

func (fs *FakeScene) TestRay(ray core.Ray) color.Color {
	fs.RecordedRays = append(fs.RecordedRays, ray)
	return fs.ColorToReturn
}
