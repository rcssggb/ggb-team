package playerclient

type sightParam struct {
	distance  float64
	direction float64
	distChng  float64
	dirChng   float64
	bodyDir   float64
	headDir   float64
}

// SightParams contains visual sensor information
type SightParams map[string]sightParam

func (s *SightParams) Init() {

}
