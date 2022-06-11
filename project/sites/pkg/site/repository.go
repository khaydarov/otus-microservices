package site

func NewRepository() Repository {
	return Repository{}
}

type Repository struct {
}

func (r *Repository) Store(site Site) error {
	return nil
}

//func (r *Repository) FindByCode()
