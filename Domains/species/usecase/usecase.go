package usecase

import (
	"challengeIskayPet/Domains/species/entity/repository"
	"challengeIskayPet/model"
)

type UsecaseInterface interface {
	Create(author model.Species) (model.Species, error)
	Delete(filfer model.QueryFilters) error
	GetOne(queryFilters model.QueryFilters) (model.Species, error)
	GetAll(filter model.QueryFilters) ([]model.Species, error)
	Update(filter *model.QueryFilters, objectToUpdate model.Species) error
}

type Usecase struct {
	repo repository.RepositoryInterface
}

func NewUsecase(repo repository.RepositoryInterface) UsecaseInterface {
	return &Usecase{
		repo: repo,
	}
}
func (u *Usecase) Create(objectToCreate model.Species) (model.Species, error) {

	errCreate := u.repo.Create(&objectToCreate)
	return objectToCreate, errCreate
}

func (u *Usecase) Delete(filter model.QueryFilters) error {

	return u.repo.Delete(&filter)
}

func (u *Usecase) GetOne(queryFilters model.QueryFilters) (model.Species, error) {

	return u.repo.GetOne(&queryFilters)
}

func (u *Usecase) GetAll(filter model.QueryFilters) ([]model.Species, error) {

	return u.repo.GetAll(filter)
}

func (u *Usecase) Update(filter *model.QueryFilters, objectToUpdate model.Species) error {

	err := u.repo.Update(filter, objectToUpdate)
	return err
}
