package repository

import (
	"challengeIskayPet/model"
)

type RepositoryInterface interface {
	Create(teacher *model.Species) error
	GetOne(queryFilter *model.QueryFilters) (model.Species, error)
	GetAll(filter model.QueryFilters) ([]model.Species, error)
	Update(filter *model.QueryFilters, user model.Species) error
	Delete(queryFilter *model.QueryFilters) error
}
