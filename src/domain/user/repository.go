package user

// Repository ...
type Repository interface {
	FindByID(ID int64) (User, error)
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
	Create(user User) error
	Update(id int64, d User) (User, error)
	Delete(id int64) error
}
