package user

func NewRepository() Repository {
	return Repository{}
}

type Repository struct {
}

func (r *Repository) Save(user User) error {
	return nil
}

func (r *Repository) FindByID(id ID) (User, error) {
	return User{}, nil
}

func (r *Repository) FindByEmailAndPassword(email Email, password string) (User, error) {
	return User{
		NewID(),
		email,
		password,
		NewType(1),
	}, nil
}
