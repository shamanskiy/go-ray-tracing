package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/utils"
)

func TestTransparent_RefactionIndexLimits(t *testing.T) {
	t.Log("When we construct a transparent material,")

	t.Log("  if we pass a refractive index less than 0.001, e.g. 0, the index is set to 0.001:")
	material := NewTransparent(0.0)
	utils.CheckResult(t, "refractive index", material.refractionIndex, core.Real(0.001))

	t.Log("  if we pass a refractive index equal or greater than 0.001, e.g. 1.5, the index is set to 1.5:")
	material = NewTransparent(1.5)
	utils.CheckResult(t, "refractive index", material.refractionIndex, core.Real(1.5))
}
