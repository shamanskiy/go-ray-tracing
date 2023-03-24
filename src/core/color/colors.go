package color

import (
	rgba "image/color"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/chewxy/math32"
	"github.com/go-gl/mathgl/mgl32"
)

var Red = NewColor(1.0, 0.0, 0.0)
var Green = NewColor(0.0, 1.0, 0.0)
var Blue = NewColor(0.0, 0.0, 1.0)
var Black = NewColor(0.0, 0.0, 0.0)
var White = NewColor(1.0, 1.0, 1.0)
var SkyBlue = NewColor(0.5, 0.7, 1.0)
var GrayMedium = NewColor(0.5, 0.5, 0.5)
var GrayLight = NewColor(0.8, 0.8, 0.8)
var Golden = NewColor(0.8, 0.6, 0.2)
var Yellow = NewColor(1.0, 1.0, 0.2)

type Color struct {
	vec mgl32.Vec3
}

func NewColor(r, g, b core.Real) Color {
	return Color{vec: mgl32.Vec3{r, g, b}}
}

func (c Color) R() core.Real {
	return c.vec.X()
}

func (c Color) G() core.Real {
	return c.vec.Y()
}

func (c Color) B() core.Real {
	return c.vec.Z()
}

func (c Color) Add(other Color) Color {
	return Color{c.vec.Add(other.vec)}
}

func (c Color) Mul(scalar core.Real) Color {
	return Color{c.vec.Mul(scalar)}
}

func (c Color) MulColor(other Color) Color {
	return NewColor(c.R()*other.R(), c.G()*other.G(), c.B()*other.B())
}

func (c Color) Div(scalar core.Real) Color {
	return NewColor(c.R()/scalar, c.G()/scalar, c.B()/scalar)
}

func (c Color) ToRGBA() rgba.RGBA {
	return rgba.RGBA{toZero255(c.R()), toZero255(c.G()), toZero255(c.B()), 255}
}

func toZero255(x core.Real) uint8 {
	return uint8(math32.Floor(255.99 * gammaCorrection(x)))
}

func gammaCorrection(input core.Real) core.Real {
	return math32.Sqrt(input)
}
