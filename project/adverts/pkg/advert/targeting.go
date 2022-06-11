package advert

func NewTargeting(devices []string, dates []string) Targeting {
	return Targeting{
		devices,
		dates,
	}
}

type Targeting struct {
	Devices []string
	Dates   []string
}
