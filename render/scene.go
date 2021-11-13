package render

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/materials"
	"github.com/Shamanskiy/go-ray-tracer/objects"
)

type Scene struct {
	objects   []objects.Object
	materials []materials.Material
}

func (s *Scene) Hit(ray core.Ray, minParam core.Real) (hit *objects.HitRecord, objectIndex int) {

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
	hit, objectIndex := s.Hit(ray, 0.0001)
	if hit != nil {
		reflection := s.materials[objectIndex].Reflect(ray, *hit)
		if reflection != nil && depth < 10 {
			nextColor := s.testRay(reflection.Ray, depth+1)
			return core.MulElem(nextColor, reflection.Attenuation)
		} else {
			return core.Color{0.0, 0.0, 0.0}
		}
	}

	unit_direction := ray.Direction.Normalize()
	t := 0.5 * (unit_direction.Y() + 1.0)
	A := core.White.Mul(1.0 - t)
	B := core.SkyBlue.Mul(t)
	return A.Add(B)
}

func (s *Scene) TestRay(ray core.Ray) core.Color {
	return s.testRay(ray, 0)
}
