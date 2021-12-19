package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
	"github.com/chewxy/math32"
)

type Transparent struct {
	refractionIndex core.Real
}

func NewTransparent(refractionIndex core.Real) Transparent {
	if refractionIndex < 1. {
		refractionIndex = 1.
	}
	return Transparent{refractionIndex: refractionIndex}
}

// Transparent materials reflect a portion of incoming light.
// The reflected portion grows with the incidence angle.
// The Schlick Law is a polynomial approximation of the reflection ratio.
func schlickLaw(cosIn core.Real, refractionIndex core.Real) core.Real {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 *= r0

	return r0 + (1-r0)*math32.Pow(1-cosIn, 5)
}

func refract(directionIn core.Vec3, normal core.Vec3, inOutRefractionRatio core.Real) (directionOut *core.Vec3, cosOut core.Real) {
	cosIn := directionIn.Dot(normal)
	cosOutSquared := 1 - (1-cosIn*cosIn)*inOutRefractionRatio*inOutRefractionRatio
	if cosOutSquared > 0 {
		inNormal := normal.Mul(cosIn)
		inTangent := directionIn.Sub(inNormal)
		outTangent := inTangent.Mul(inOutRefractionRatio)
		cosOut := math32.Sqrt(cosOutSquared)
		outNormal := inNormal.Mul(cosOut / math32.Abs(cosIn))
		directionOut := outTangent.Add(outNormal)
		return &directionOut, cosOut
	} else {
		// The ray enters a less dense material. The incidence angle is too large, so the out angle is > 90deg.
		// There is no refraction, instead we have a full reflection.
		return nil, 0
	}
}

func computeReflectionRatio(rayComesFromOutside bool, cosIn core.Real, cosOut core.Real, refractionIndex core.Real) core.Real {
	if rayComesFromOutside {
		return schlickLaw(math32.Abs(cosIn), refractionIndex)
	} else {
		return schlickLaw(cosOut, refractionIndex)
	}
}

func (m Transparent) Reflect(ray core.Ray, hit objects.HitRecord) *Reflection {
	directionIn := ray.Direction.Normalize()
	cosIn := directionIn.Dot(hit.Normal)

	var inOutRefractionRatio core.Real
	rayComesFromOutside := cosIn <= 0.
	if rayComesFromOutside {
		inOutRefractionRatio = 1. / m.refractionIndex
	} else {
		inOutRefractionRatio = m.refractionIndex
	}

	refractedDirection, cosOut := refract(directionIn, hit.Normal, inOutRefractionRatio)
	reflectedDirection := core.Reflect(directionIn, hit.Normal)
	if refractedDirection != nil {
		reflectionRatio := computeReflectionRatio(rayComesFromOutside, cosIn, cosOut, m.refractionIndex)
		if core.Random().From01() < reflectionRatio {
			reflectedRay := core.Ray{hit.Point, reflectedDirection}
			return &Reflection{reflectedRay, core.White}
		} else {

			refractedRay := core.Ray{hit.Point, *refractedDirection}
			// transparent material doesn't alter the reflection color
			return &Reflection{refractedRay, core.White}
		}

	} else {
		reflectedRay := core.Ray{hit.Point, reflectedDirection}
		return &Reflection{reflectedRay, core.White}
	}
}
