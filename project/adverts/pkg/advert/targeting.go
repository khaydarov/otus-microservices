package advert

func NewTargeting(devices, dates []string, hits, cost int) Targeting {
	return Targeting{
		devices,
		dates,
		hits,
		cost,
	}
}

type Targeting struct {
	Devices []string
	Dates   []string
	Hits    int
	Cost    int
}
