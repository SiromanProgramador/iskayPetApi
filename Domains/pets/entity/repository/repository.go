package repository

import (
	"challengeIskayPet/model"
	pb "iskayPetMicro/api"
)

//define repository Interface
type RepositoryInterface interface {
	Create(pet pb.Pet) (*pb.Pet, error)
	GetStatistics(queryFilter *model.QueryFilters) (*pb.ResponseStatistics, error)
	GetAll(filter model.QueryFilters) ([]*pb.Pet, error)
}
