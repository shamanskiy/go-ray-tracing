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
	if refractionIndex < 1 {
		refractionIndex = 1
	}
	return Transparent{refractionIndex: refractionIndex}
}

// Transparent materials reflect a portion of incoming light.
// The reflected portion grows with the incidence angle.
// The Schlick Law is a polynomial approximation of the reflection ratio.
/*func schlickLaw(incidenceAngle core.Real, refractionIndex core.Real) core.Real {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	probability *= probability

	return probability + (1-probability)*math32.Pow(1-incidenceAngle, 5)
}*/

// The normal must point to the side from where the incident ray comes.
func refract(directionIn core.Vec3, normal core.Vec3, inOutRefractionRatio core.Real) *core.Vec3 {
	cosIn := directionIn.Dot(normal)
	cosOutSquared := 1 - (1-cosIn*cosIn)*inOutRefractionRatio*inOutRefractionRatio
	if cosOutSquared > 0 {
		inNormal := directionIn.Sub(normal.Mul(cosIn))
		inTangent := directionIn.Sub(inNormal)
		outTangent := inTangent.Mul(inOutRefractionRatio)
		outNormal := normal.Mul(-math32.Sqrt(cosOutSquared))
		outDir := outTangent.Add(outNormal)
		return &outDir
	} else {
		// The ray enters a less dense material. The incidence angle is too large, so the out angle is > 90deg.
		// There is no refraction, instead we have a full reflection.
		return nil
	}
}

func (m Transparent) Reflect(ray core.Ray, hit objects.HitRecord) *Reflection {
	directionIn := ray.Direction.Normalize()
	incidenceCos := directionIn.Dot(hit.Normal)

	var inOutRefractionRatio core.Real
	var refractionNormal core.Vec3
	if incidenceCos > 0 {
		// ray comes from inside the material
		inOutRefractionRatio = m.refractionIndex
		refractionNormal = hit.Normal.Mul(-1)
	} else {
		// ray comes from outside the material
		inOutRefractionRatio = 1. / m.refractionIndex
		refractionNormal = hit.Normal
	}

	refractedDirection := refract(directionIn, refractionNormal, inOutRefractionRatio)
	if refractedDirection != nil {
		refractedRay := core.Ray{hit.Point, *refractedDirection}
		// transparent material doesn't alter the reflection color
		return &Reflection{refractedRay, core.White}
	} else {
		reflectedDirection := core.Reflect(directionIn, hit.Normal)
		reflectedRay := core.Ray{hit.Point, reflectedDirection}
		return &Reflection{reflectedRay, core.White}
	}
}
