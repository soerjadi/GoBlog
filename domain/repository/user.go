package repository

import "github.com/soerjadi/GoBlog/domain"

// UserRepository represent repositorf of user
type UserRepository interface {
	GetByID(id int64) (*domain.User, error)
	GetList(offset int, limit int) ([]*domain.User, error)
	Save(*domain.User) error
	Delete(id int64) error
}
