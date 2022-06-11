package advert

func NewAdvert(title, description, link, image string, devices, dates []string) Advert {
	return Advert{
		NewID(),
		title,
		description,
		link,
		image,
		NewTargeting(devices, dates),
	}
}

type Advert struct {
	ID          ID
	Title       string
	Description string
	Link        string
	Image       string
	Targeting   Targeting
}
