package usecase

import (
	"challengeIskayPet/Domains/pets/entity/repository"
	"challengeIskayPet/model"
	pb "iskayPetMicro/api"
)

type UsecaseInterface interface {
	Create(pet pb.Pet) (*pb.Pet, error)
	GetStatistics(queryFilters model.QueryFilters) (*pb.ResponseStatistics, error)
	GetAll(filter model.QueryFilters) ([]*pb.Pet, error)
}

type Usecase struct {
	repo repository.RepositoryInterface
}

func NewUsecase(repo repository.RepositoryInterface) UsecaseInterface {
	return &Usecase{
		repo: repo,
	}
}
func (u *Usecase) Create(objectToCreate pb.Pet) (*pb.Pet, error) {

	pet, errCreate := u.repo.Create(objectToCreate)
	return pet, errCreate
}

func (u *Usecase) GetStatistics(queryFilters model.QueryFilters) (*pb.ResponseStatistics, error) {

	response, errResponse := u.repo.GetStatistics(&queryFilters)
	return response, errResponse
}

func (u *Usecase) GetAll(filter model.QueryFilters) ([]*pb.Pet, error) {
	response, errResponse := u.repo.GetAll(filter)
	return response, errResponse
}
