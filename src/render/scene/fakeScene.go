package scene

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
)

type FakeScene struct{}

func (fs *FakeScene) TestRay(ray core.Ray) color.Color {
	return color.Red
}
