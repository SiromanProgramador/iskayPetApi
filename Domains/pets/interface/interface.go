package _interface

import (
	Usecase "challengeIskayPet/Domains/pets/usecase"
	"challengeIskayPet/model"
	"challengeIskayPet/presenters"
	"encoding/json"
	"io/ioutil"
	pb "iskayPetMicro/api"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type InterfaceInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetStatistics(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type Interface struct {
	usecase Usecase.UsecaseInterface
}

func PetsInterface(usecase Usecase.UsecaseInterface) InterfaceInterface {
	return &Interface{
		usecase: usecase,
	}
}

func (ui *Interface) GetAll(w http.ResponseWriter, r *http.Request) {

	var qfilter model.QueryFilters
	var response []interface{}

	authors, errAuthors := ui.usecase.GetAll(qfilter)

	if errAuthors != nil {
		presenters.ReturnHttpError(errAuthors, w, http.StatusBadRequest)
		return
	}
	err := presenters.ArrayStructToBson(authors, &response)
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

	response, err := presenters.StructToBson(objectToReturn)
	if err != nil {
		presenters.ReturnHttpError(err, w, http.StatusBadRequest)
		return
	}

	presenters.ReturnHttpPayload(response, w)
}

func (ui *Interface) Create(w http.ResponseWriter, r *http.Request) {

	//GET body update request
	var objectToCreate pb.Pet
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

	_, errCreate := ui.usecase.Create(objectToCreate)
	if errCreate != nil {
		presenters.ReturnHttpError(errCreate, w, http.StatusInternalServerError)
		return
	}

	presenters.ReturnHttpPayload(bson.M{"success": true}, w)
}
