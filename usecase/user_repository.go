package usecase

import "app/domain"

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindById(id int) (domain.User, error)
}
