package core

type Scene struct {
	objects []Object
}

func (s *Scene) Hit(ray Ray) HitRecord {
	var hit HitRecord

	for _, obj := range s.objects {
		tempHit := obj.Hit(ray)
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
