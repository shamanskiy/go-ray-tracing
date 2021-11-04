package core

type Scene struct {
	objects []Object
}

func (s *Scene) Hit(ray Ray) HitRecord {
	return s.HitWithMin(ray, 0.0)
}

func (s *Scene) HitWithMin(ray Ray, minParam Real) HitRecord {
	var hit HitRecord

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

func (s *Scene) Add(obj ...Object) {
	s.objects = append(s.objects, obj...)
}
