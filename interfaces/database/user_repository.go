package database

import (
	"app/domain"
)

type UserRepository struct {
	SqlHandler
}

func (userRepository *UserRepository) FindAll() (users domain.Users, err error) {
	if err = userRepository.Find(&users).Error; err != nil {
		return
	}
	return
}
