package calc

import "testing"

func TestSphere(t *testing.T) {
	expDiam := 4.0
	expSurfaceArea := 50.26548246
	expVol := 33.51032164

	s := Sphere{2.0}
	actDiam := s.Diam()
	actSurfaceArea := s.SurfaceArea()
	actVol := s.Vol()

	if !Near(expDiam, actDiam) {
		t.Error("Expected", expDiam, "got", actDiam)
	}
	if !Near(expSurfaceArea, actSurfaceArea) {
		t.Error("Expected", expSurfaceArea, "got", actSurfaceArea)
	}
	if !Near(expVol, actVol) {
		t.Error("Expected", expVol, "got", actVol)
	}
}
