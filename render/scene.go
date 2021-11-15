package render

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/materials"
	"github.com/Shamanskiy/go-ray-tracer/objects"
	"github.com/chewxy/math32"
)

type Scene struct {
	objects        []objects.Object
	materials      []materials.Material
	SkyColorTop    core.Color
	SkyColorBottom core.Color
}

func (s *Scene) hitClosestObject(ray core.Ray, minParam core.Real) (hit *objects.HitRecord, objectIndex int) {

	for i := range s.objects {
		tempHit := s.objects[i].HitWithMin(ray, minParam)
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

func (s *Scene) testRay(ray core.Ray, depth int) core.Color {
	hit, objectIndex := s.hitClosestObject(ray, 0.0001)
	if hit != nil {
		reflection := s.materials[objectIndex].Reflect(ray, *hit)
		if reflection != nil && depth < 10 {
			reflectedRayColor := s.testRay(reflection.Ray, depth+1)
			return core.MulElem(reflectedRayColor, reflection.Attenuation)
		} else {
			return core.Black
		}
	}

	return s.computeSkyColor(ray)
}

func (s *Scene) TestRay(ray core.Ray) core.Color {
	return s.testRay(ray, 0)
}

func (s *Scene) computeSkyColor(ray core.Ray) core.Color {
	unit_direction := ray.Direction.Normalize()

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
