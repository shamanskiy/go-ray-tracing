package color

import (
	rgba "image/color"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/chewxy/math32"
	"github.com/go-gl/mathgl/mgl32"
)

var Red = New(1.0, 0.0, 0.0)
var Green = New(0.0, 1.0, 0.0)
var Blue = New(0.0, 0.0, 1.0)
var Black = New(0.0, 0.0, 0.0)
var White = New(1.0, 1.0, 1.0)
var SkyBlue = New(0.5, 0.7, 1.0)
var GrayMedium = New(0.5, 0.5, 0.5)
var GrayLight = New(0.8, 0.8, 0.8)
var Golden = New(0.8, 0.6, 0.2)
var Yellow = New(1.0, 1.0, 0.2)

type Color struct {
	vec mgl32.Vec3
}

func New(r, g, b core.Real) Color {
	return Color{vec: mgl32.Vec3{r, g, b}}
}

func FromVec3(vec core.Vec3) Color {
	return New(vec.X(), vec.Y(), vec.Z())
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
	return New(c.R()*other.R(), c.G()*other.G(), c.B()*other.B())
}

func (c Color) Div(scalar core.Real) Color {
	return New(c.R()/scalar, c.G()/scalar, c.B()/scalar)
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

func Interpolate(A, B Color, t core.Real) Color {
	return A.Mul(1 - t).Add(B.Mul(t))
}
