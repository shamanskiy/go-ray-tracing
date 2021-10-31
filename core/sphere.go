package core

type Sphere struct {
	Center Vec3
	Radius Real
}

func (s Sphere) Hit(ray Ray) HitRecord {
	centerToOrigin := ray.Origin.Sub(s.Center)

	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(centerToOrigin)
	c := centerToOrigin.Dot(centerToOrigin) - s.Radius*s.Radius
	left, right, err := solveQuadraticEquation(a, b, c)

	var hit HitRecord
	if err != nil {
		return hit
	}

	if left >= ray.MinParam {
		hit.Param = left
		hit.Hit = true
	}

	if !hit.Hit && right >= ray.MinParam {
		hit.Param = right
		hit.Hit = true
	}

	if hit.Hit {
		hit.Point = ray.Eval(hit.Param)
		hit.Normal = hit.Point.Sub(s.Center).Mul(1 / s.Radius)
	}

	return hit
}
