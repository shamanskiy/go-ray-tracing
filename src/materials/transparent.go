package materials

import (
	"fmt"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/chewxy/math32"
)

type Transparent struct {
	refractor  RefractionCalculator
	randomizer random.RandomGenerator
}

func NewTransparent(refractionIndex core.Real, randomizer random.RandomGenerator) Transparent {
	if refractionIndex < 1 {
		panic(fmt.Errorf("refractionIndex must be at least 1, got %f", refractionIndex))
	}
	return Transparent{refractor: RefractionCalculator{refractionIndex}, randomizer: randomizer}
}

func (m Transparent) Reflect(incidentDirection, hitPoint, normalAtHitPoint core.Vec3) *Reflection {
	refractedDirection, reflectionRatio := m.refractor.Refract(incidentDirection, normalAtHitPoint)
	reflectedDirection := incidentDirection.Reflect(normalAtHitPoint)

	// Transparent material reflects a portion of the incoming light
	if refractedDirection != nil && m.randomizer.Real() > reflectionRatio {
		refractedRay := core.NewRay(hitPoint, *refractedDirection)
		return &Reflection{refractedRay, color.White}
	} else {
		// Full internal reflection
		reflectedRay := core.NewRay(hitPoint, reflectedDirection)
		return &Reflection{reflectedRay, color.White}
	}
}

type RefractionCalculator struct {
	RefractionIndex core.Real
}

// Returns the direction of the refracted ray and the portion of the reflected light.
// If the ray comes from material at a too large incident angle, returns nil (full internal reflection).
func (r RefractionCalculator) Refract(dirIn core.Vec3, normal core.Vec3) (dirOut *core.Vec3, reflectionRatio core.Real) {
	inDotNormal := dirIn.Dot(normal)
	rayFromOutside := inDotNormal <= 0
	inOutRefractionRatio := r.computeInOutRefractionRatio(rayFromOutside)

	cosIn := inDotNormal / dirIn.Len()
	cosOutSquared := 1 - (1-cosIn*cosIn)*inOutRefractionRatio*inOutRefractionRatio

	if cosOutSquared > 0 {
		inNormal := normal.Mul(inDotNormal)
		inTangent := dirIn.Sub(inNormal)
		outTangent := inTangent.Mul(inOutRefractionRatio)
		cosOut := math32.Sqrt(cosOutSquared)
		outNormal := inNormal.Mul(cosOut / math32.Abs(cosIn))
		directionOut := outTangent.Add(outNormal)

		return &directionOut, r.computeReflectionRatio(rayFromOutside, cosIn, cosOut)
	} else {
		// The ray enters a less dense material. The incidence angle is too large, so the out angle is > 90deg.
		// There is no refraction, instead we have a full reflection.
		return nil, 1.
	}
}

func (r RefractionCalculator) computeInOutRefractionRatio(rayFromOutside bool) core.Real {
	if rayFromOutside {
		return 1. / r.RefractionIndex
	} else {
		return r.RefractionIndex
	}
}

func (r RefractionCalculator) computeReflectionRatio(rayFromOutside bool, cosIn core.Real, cosOut core.Real) core.Real {
	if rayFromOutside {
		return r.SchlickLaw(math32.Abs(cosIn))
	} else {
		return r.SchlickLaw(cosOut)
	}
}

// Transparent materials reflect a portion of incoming light.
// The reflected portion grows with the incidence angle.
// The Schlick Law is a polynomial approximation of the reflection ratio.
func (r RefractionCalculator) SchlickLaw(cosIn core.Real) core.Real {
	r0 := (1 - r.RefractionIndex) / (1 + r.RefractionIndex)
	r0 *= r0

	return r0 + (1-r0)*math32.Pow(1-cosIn, 5)
}
