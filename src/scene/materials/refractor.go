package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/chewxy/math32"
)

type Refraction struct {
	Direction       *core.Vec3
	ReflectionRatio core.Real
}

func (refaction Refraction) FullInternalReflection() bool {
	return refaction.Direction == nil
}

type RefractionCalculator struct {
	RefractionIndex core.Real
}

func NewRefractionCalculator(refractionIndex core.Real) RefractionCalculator {
	return RefractionCalculator{refractionIndex}
}

// Returns the direction of the refracted ray and the portion of the reflected light.
// If the ray comes from material at a too large incident angle, returns nil (full internal reflection).
func (r RefractionCalculator) Refract(incidentDir core.Vec3, normal core.Vec3) Refraction {
	fromOutside := comesFromOutside(incidentDir, normal)
	if fromOutside {
		return r.refractEnterMaterial(incidentDir, normal)
	} else {
		return r.refractExitMaterial(incidentDir, normal)
	}
}

func comesFromOutside(incidentDirection core.Vec3, normal core.Vec3) bool {
	return incidentDirection.Dot(normal) <= 0
}

// Transparent materials reflect a portion of incoming light, the portion grows with the incidence angle.
// The Schlick Law is a polynomial approximation of the reflection ratio.
func (r RefractionCalculator) schlickLaw(cosOutsideMaterial core.Real) core.Real {
	r0 := (1 - r.RefractionIndex) / (1 + r.RefractionIndex)
	r0 *= r0
	return r0 + (1-r0)*math32.Pow(1-cosOutsideMaterial, 5)
}

func (r RefractionCalculator) refractEnterMaterial(inDir core.Vec3, normal core.Vec3) Refraction {
	inDotNormal := inDir.Dot(normal)
	cosIn := math32.Abs(inDotNormal / inDir.Len())
	cosOut := math32.Sqrt(1 - (1-cosIn*cosIn)/(r.RefractionIndex*r.RefractionIndex))

	inNormalComponent := normal.Mul(inDotNormal)
	inTangentComponent := inDir.Sub(inNormalComponent)

	outTangentComponent := inTangentComponent.Div(r.RefractionIndex)
	outNormalComponent := inNormalComponent.Mul(cosOut / cosIn)

	directionOut := outTangentComponent.Add(outNormalComponent)
	reflectionRatio := r.schlickLaw(cosIn)

	return Refraction{&directionOut, reflectionRatio}
}

func (r RefractionCalculator) refractExitMaterial(inDir core.Vec3, normal core.Vec3) Refraction {
	inDotNormal := inDir.Dot(normal)
	cosIn := inDotNormal / inDir.Len()
	cosOutSquared := 1 - (1-cosIn*cosIn)*r.RefractionIndex*r.RefractionIndex

	if cosOutSquared <= 0 {
		// Incident angle too large, full internal reflection
		return Refraction{nil, 1.}
	}

	cosOut := math32.Sqrt(cosOutSquared)
	inNormalComponent := normal.Mul(inDotNormal)
	inTangentComponent := inDir.Sub(inNormalComponent)

	outTangentComponent := inTangentComponent.Mul(r.RefractionIndex)
	outNormalComponent := inNormalComponent.Mul(cosOut / cosIn)

	directionOut := outTangentComponent.Add(outNormalComponent)
	reflectionRatio := r.schlickLaw(cosOut)

	return Refraction{&directionOut, reflectionRatio}
}
