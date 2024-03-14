package primitives

import "testing"

func TestPrimitiveInterfaces(t *testing.T) {
	var _ IAny = new(Boolean)
	var _ IAny = new(Character)
	var _ IAny = new(Double)
	var _ IAny = new(Real)
	var _ IAny = new(Integer)
	var _ IAny = new(Integer64)
	var _ IAny = new(Octet)
	var _ IAny = new(String)
	var _ IAny = new(Uri)
	var _ IOrdered = new(Character)
	var _ IOrdered = new(Octet)
	var _ IOrdered = new(String)
	var _ IOrdered = new(Uri)
	var _ INumeric = new(Double)
	var _ INumeric = new(Real)
	var _ INumeric = new(Integer)
	var _ INumeric = new(Integer64)
}
