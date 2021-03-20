package usecase

import "app/domain"

type UserRepository interface {
	FindAll() (domain.Users, error)
}
