package calc

import "testing"

func TestSphere(t *testing.T) {
	expDiam := 4.0
	expSurfArea := 50.26548245743669
	expVol := 33.51032163829113

	s := Sphere{2.0}
	actDiam := s.Diam()
	actSurfArea := s.SurfArea()
	actVol := s.Vol()

	if !Near(expDiam, actDiam) {
		t.Error("Expected", expDiam, "got", actDiam)
	}
	if !Near(expSurfArea, actSurfArea) {
		t.Error("Expected", expSurfArea, "got", actSurfArea)
	}
	if !Near(expVol, actVol) {
		t.Error("Expected", expVol, "got", actVol)
	}
}
