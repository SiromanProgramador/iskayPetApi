package _interface

import (
	Usecase "challengeIskayPet/Domains/pets/usecase"
	"challengeIskayPet/model"
	"challengeIskayPet/presenters"
	"encoding/json"
	"io/ioutil"
	pb "iskayPetMicro/api"
	"net/http"
)

//define interface to interface layer
type InterfaceInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetStatistics(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

//define struct to interface layer
type Interface struct {
	usecase Usecase.UsecaseInterface
}

//constructor to petInterface that return a pointer Interface with next layer Interface (usecaseInterface)
func PetsInterface(usecase Usecase.UsecaseInterface) InterfaceInterface {
	return &Interface{
		usecase: usecase,
	}
}

func (ui *Interface) GetAll(w http.ResponseWriter, r *http.Request) {

	var qfilter model.QueryFilters
	var response []interface{}

	pets, errPets := ui.usecase.GetAll(qfilter)

	if errPets != nil {
		presenters.ReturnHttpError(errPets, w, http.StatusBadRequest)
		return
	}

	//parse struct to bson
	err := presenters.ArrayStructToBson(pets, &response)
	if err != nil {
		presenters.ReturnHttpError(err, w, http.StatusBadRequest)
		return
	}

	presenters.ReturnHttpPayload(response, w)
}

func (ui *Interface) GetStatistics(w http.ResponseWriter, r *http.Request) {

	//SET filter
	var qfilter model.QueryFilters

	objectToReturn, errorFind := ui.usecase.GetStatistics(qfilter)
	if errorFind != nil {
		presenters.ReturnHttpError(errorFind, w, http.StatusNotFound)
		return
	}

	//parse struct to bson
	response, err := presenters.StructToBson(objectToReturn)
	if err != nil {
		presenters.ReturnHttpError(err, w, http.StatusBadRequest)
		return
	}

	presenters.ReturnHttpPayload(response, w)
}

func (ui *Interface) Create(w http.ResponseWriter, r *http.Request) {

	var objectToCreate pb.Pet

	//GET body update request
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		presenters.ReturnHttpError(err, w, http.StatusInternalServerError)
		return
	}

	errUnmarshal := json.Unmarshal(body, &objectToCreate)
	if errUnmarshal != nil {
		presenters.ReturnHttpError(errUnmarshal, w, http.StatusInternalServerError)
		return
	}

	pet, errCreate := ui.usecase.Create(objectToCreate)
	if errCreate != nil {
		presenters.ReturnHttpError(errCreate, w, http.StatusInternalServerError)
		return
	}

	//parse struct to bson
	response, err := presenters.StructToBson(pet)
	if err != nil {
		presenters.ReturnHttpError(err, w, http.StatusBadRequest)
		return
	}

	presenters.ReturnHttpPayload(response, w)
}
