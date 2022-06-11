package advert

func NewRepository() Repository {
	return Repository{}
}

type Repository struct {
}

func (r *Repository) FindByID(id ID) (Advert, error) {
	return Advert{}, nil
}

func (r *Repository) Store(advert Advert) error {
	return nil
}
