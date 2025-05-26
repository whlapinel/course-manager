package user

type Repository interface {
	ByID(id string) (User, error)
	Save(user User) (string, error) // returns id
	All() ([]User, error)
}
