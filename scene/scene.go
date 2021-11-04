package scene

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
)

type Scene struct {
	objects []objects.Object
}

func (s *Scene) Hit(ray core.Ray) core.HitRecord {
	return s.HitWithMin(ray, 0.0)
}

func (s *Scene) HitWithMin(ray core.Ray, minParam core.Real) core.HitRecord {
	var hit core.HitRecord

	for i := range s.objects {
		tempHit := s.objects[i].HitWithMin(ray, minParam)
		if tempHit.Hit {
			if !hit.Hit || (hit.Hit && hit.Param > tempHit.Param) {
				hit = tempHit
			}
		}
	}

	return hit
}

func (s *Scene) Add(obj ...objects.Object) {
	s.objects = append(s.objects, obj...)
}
