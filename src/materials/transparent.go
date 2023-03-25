package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/chewxy/math32"
)

type Transparent struct {
	RefractionIndex core.Real
	randomizer      random.RandomGenerator
}

func NewTransparent(refractionIndex core.Real, randomizer random.RandomGenerator) Transparent {
	if refractionIndex < 1. {
		refractionIndex = 1.
	}
	return Transparent{RefractionIndex: refractionIndex, randomizer: randomizer}
}

// Transparent materials reflect a portion of incoming light.
// The reflected portion grows with the incidence angle.
// The Schlick Law is a polynomial approximation of the reflection ratio.
func SchlickLaw(cosIn core.Real, refractionIndex core.Real) core.Real {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 *= r0

	return r0 + (1-r0)*math32.Pow(1-cosIn, 5)
}

func computeInOutRefractionRatio(rayFromOutside bool, refractionIndex core.Real) core.Real {
	if rayFromOutside {
		return 1. / refractionIndex
	} else {
		return refractionIndex
	}
}

// Returns the direction of the refracted ray and the portion of the reflected light.
// If the ray comes from material at a too large incident angle, returns nil (full internal reflection).
func ComputeRefraction(dirIn core.Vec3, normal core.Vec3, refractionIndex core.Real) (dirOut *core.Vec3, reflectionRatio core.Real) {
	inDotNormal := dirIn.Dot(normal)
	rayFromOutside := inDotNormal <= 0
	inOutRefractionRatio := computeInOutRefractionRatio(rayFromOutside, refractionIndex)

	cosIn := inDotNormal / dirIn.Len()
	cosOutSquared := 1 - (1-cosIn*cosIn)*inOutRefractionRatio*inOutRefractionRatio

	if cosOutSquared > 0 {
		inNormal := normal.Mul(inDotNormal)
		inTangent := dirIn.Sub(inNormal)
		outTangent := inTangent.Mul(inOutRefractionRatio)
		cosOut := math32.Sqrt(cosOutSquared)
		outNormal := inNormal.Mul(cosOut / math32.Abs(cosIn))
		directionOut := outTangent.Add(outNormal)

		return &directionOut, computeReflectionRatio(rayFromOutside, cosIn, cosOut, refractionIndex)
	} else {
		// The ray enters a less dense material. The incidence angle is too large, so the out angle is > 90deg.
		// There is no refraction, instead we have a full reflection.
		return nil, 1.
	}
}

func computeReflectionRatio(rayFromOutside bool, cosIn core.Real, cosOut core.Real, refractionIndex core.Real) core.Real {
	if rayFromOutside {
		return SchlickLaw(math32.Abs(cosIn), refractionIndex)
	} else {
		return SchlickLaw(cosOut, refractionIndex)
	}
}

func (m Transparent) Reflect(incidentDirection, hitPoint, normalAtHitPoint core.Vec3) *Reflection {
	refractedDirection, reflectionRatio := ComputeRefraction(incidentDirection, normalAtHitPoint, m.RefractionIndex)
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
