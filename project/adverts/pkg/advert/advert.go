package advert

func NewAdvert(userID string, title, description, link, image string, devices, dates []string, hits, cost int) Advert {
	return Advert{
		NewID(),
		userID,
		title,
		description,
		link,
		image,
		NewTargeting(devices, dates, hits, cost),
	}
}

type Advert struct {
	ID          ID
	UserID      string
	Title       string
	Description string
	Link        string
	Image       string
	Targeting   Targeting
}
