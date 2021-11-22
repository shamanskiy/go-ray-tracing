package materials

import "github.com/Shamanskiy/go-ray-tracer/core"

type Transparent struct {
	refractionIndex core.Real
}

func NewTransparent(refractionIndex core.Real) Transparent {
	if refractionIndex < 0.001 {
		refractionIndex = 0.001
	}
	return Transparent{refractionIndex: refractionIndex}
}
