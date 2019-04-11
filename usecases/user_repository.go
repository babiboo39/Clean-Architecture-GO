package usecases

import "modul1/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
	FindByEmail(email string, password string) (domain.User, error)
}
