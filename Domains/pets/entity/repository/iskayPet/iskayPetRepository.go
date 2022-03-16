package mongodb

import (
	"challengeIskayPet/Domains/pets/entity/repository"
	"challengeIskayPet/model"
	"context"
	pb "iskayPetMicro/api"
	"log"
	"time"

	"google.golang.org/grpc"
	mgo "gopkg.in/mgo.v2"
)

type repo struct {
	session *mgo.Session
}

//IskayPet Microservice repository
//in this layer we have all connections to our microservice

func NewIskayPetRepository(session *mgo.Session) repository.RepositoryInterface {
	return &repo{
		session: session,
	}
}

func (r *repo) Create(objectToCreate pb.Pet) (*pb.Pet, error) {

	//define connection with grpc
	conn, err := grpc.Dial("localhost:7770", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCreatePetServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//call createPet Method
	response, errCreate := c.CreatePet(ctx, &pb.CreatePetRequest{Pet: &objectToCreate})

	return response.Pet, errCreate
}

func (r *repo) GetStatistics(queryFilter *model.QueryFilters) (*pb.ResponseStatistics, error) {

	//define connection with grpc
	conn, err := grpc.Dial("localhost:7770", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGetStatisticsServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//call get Statistics method, the function is prepare to be able to send a particular petName to get his statistic
	response, errGetStatistics := c.GetStatistics(ctx, &pb.GetStatisticsRequest{PetName: ""})

	return response.Statistics, errGetStatistics
}

func (r *repo) GetAll(filter model.QueryFilters) ([]*pb.Pet, error) {

	//define connection with grpc
	conn, err := grpc.Dial("localhost:7770", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGetPetsServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//call GetAlPets Method, Funciton is prepare to be able to send a filter
	pets, errPets := c.GetPets(ctx, &pb.GetPetsRequest{Filter: ""})

	return pets.Pets, errPets
}
