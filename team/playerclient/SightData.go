package playerclient

type sightData struct {
	time      int64
	distance  float64
	direction float64
	distChng  float64
	dirChng   float64
	bodyDir   float64
	headDir   float64
}

// SightMap contains visual sensor information
type SightMap map[string]*sightData

// Init allocates SightMap
func (s *SightMap) Init() {
	sightMap := make(SightMap)
	// Initialize flag markers
	sightMap["f c"] = &sightData{}
	sightMap["f c t"] = &sightData{}
	sightMap["f c b"] = &sightData{}
	sightMap["g l"] = &sightData{}
	sightMap["f g l t"] = &sightData{}
	sightMap["f g l b"] = &sightData{}
	sightMap["g r"] = &sightData{}
	sightMap["f g r t"] = &sightData{}
	sightMap["f g r b"] = &sightData{}
	sightMap["f p l c"] = &sightData{}
	sightMap["f p l t"] = &sightData{}
	sightMap["f p l b"] = &sightData{}
	sightMap["f p r c"] = &sightData{}
	sightMap["f p r t"] = &sightData{}
	sightMap["f p r b"] = &sightData{}
	sightMap["f l t"] = &sightData{}
	sightMap["f l b"] = &sightData{}
	sightMap["f r t"] = &sightData{}
	sightMap["f r b"] = &sightData{}
	sightMap["f t l 50"] = &sightData{}
	sightMap["f t l 40"] = &sightData{}
	sightMap["f t l 30"] = &sightData{}
	sightMap["f t l 20"] = &sightData{}
	sightMap["f t l 10"] = &sightData{}
	sightMap["f t 0"] = &sightData{}
	sightMap["f t r 50"] = &sightData{}
	sightMap["f t r 40"] = &sightData{}
	sightMap["f t r 30"] = &sightData{}
	sightMap["f t r 20"] = &sightData{}
	sightMap["f t r 10"] = &sightData{}
	sightMap["f b l 50"] = &sightData{}
	sightMap["f b l 40"] = &sightData{}
	sightMap["f b l 30"] = &sightData{}
	sightMap["f b l 20"] = &sightData{}
	sightMap["f b l 10"] = &sightData{}
	sightMap["f b 0"] = &sightData{}
	sightMap["f b r 50"] = &sightData{}
	sightMap["f b r 40"] = &sightData{}
	sightMap["f b r 30"] = &sightData{}
	sightMap["f b r 20"] = &sightData{}
	sightMap["f b r 10"] = &sightData{}
	sightMap["f l t 30"] = &sightData{}
	sightMap["f l t 20"] = &sightData{}
	sightMap["f l t 10"] = &sightData{}
	sightMap["f l 0"] = &sightData{}
	sightMap["f l b 30"] = &sightData{}
	sightMap["f l b 20"] = &sightData{}
	sightMap["f l b 10"] = &sightData{}
	sightMap["f r t 30"] = &sightData{}
	sightMap["f r t 20"] = &sightData{}
	sightMap["f r t 10"] = &sightData{}
	sightMap["f r 0"] = &sightData{}
	sightMap["f r b 30"] = &sightData{}
	sightMap["f r b 20"] = &sightData{}
	sightMap["f r b 10"] = &sightData{}

	*s = sightMap
}
