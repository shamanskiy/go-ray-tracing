package core

type Camera struct {
	Origin            Vec3
	Upper_left_corner Vec3
	Horizontal        Vec3
	Vertical          Vec3
}

/*func NewCamera() camera {
	return camera{
		origin:            Vec3{0.0, 0.0, 0.0},
		upper_left_corner: Vec3{-2.0, 1.0, -1.0},
		horizontal:        Vec3{4.0, 0.0, 0.0},
		vertical:          Vec3{0.0, -2.0, 0.0}}
}*/

func (c *Camera) GetRay(u, v Real) Ray {
	ray := Ray{
		Origin:    c.Origin,
		Direction: c.Upper_left_corner.Add(c.Horizontal.Mul(u)).Add(c.Vertical.Mul(v))}

	return ray
}
