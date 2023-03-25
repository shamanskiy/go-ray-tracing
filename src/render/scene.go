package render

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/chewxy/math32"
)

type Scene struct {
	objects        []objects.Object
	materials      []materials.Material
	SkyColorTop    color.Color
	SkyColorBottom color.Color
}

func (s *Scene) HitClosestObject(ray core.Ray, minParam core.Real) (hit *objects.HitRecord, objectIndex int) {

	for i := range s.objects {
		tempHit := s.objects[i].Hit(ray, minParam)
		if tempHit != nil {
			if hit == nil || (hit != nil && hit.Param > tempHit.Param) {
				hit = tempHit
				objectIndex = i
			}
		}
	}

	return hit, objectIndex
}

func (s *Scene) Add(object objects.Object, material materials.Material) {
	s.objects = append(s.objects, object)
	s.materials = append(s.materials, material)
}

func (s *Scene) testRay(ray core.Ray, depth int) color.Color {
	hit, objectIndex := s.HitClosestObject(ray, 0.0001)
	if hit != nil {
		reflection := s.materials[objectIndex].Reflect(ray, *hit)
		if reflection != nil && depth < 10 {
			reflectedRayColor := s.testRay(reflection.Ray, depth+1)
			return reflectedRayColor.MulColor(reflection.Attenuation)
		} else {
			return color.Black
		}
	}

	return s.computeSkyColor(ray)
}

func (s *Scene) TestRay(ray core.Ray) color.Color {
	return s.testRay(ray, 0)
}

func (s *Scene) computeSkyColor(ray core.Ray) color.Color {
	unit_direction := ray.Direction().Normalize()

	var t float32
	// check for zero length vector
	if math32.IsNaN(unit_direction.X()) {
		t = 0.5
	} else {
		t = 0.5 * (unit_direction.Y() + 1.0)
	}

	A := s.SkyColorBottom.Mul(1.0 - t)
	B := s.SkyColorTop.Mul(t)
	return A.Add(B)
}
