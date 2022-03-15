package _interface

import (
	Usecase "challengeIskayPet/Domains/species/usecase"
	"challengeIskayPet/model"
	"challengeIskayPet/presenters"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type InterfaceInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type Interface struct {
	usecase Usecase.UsecaseInterface
}

func SpeciesInterface(usecase Usecase.UsecaseInterface) InterfaceInterface {
	return &Interface{
		usecase: usecase,
	}
}

func (ui *Interface) GetAll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("entra")
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

func (ui *Interface) GetOne(w http.ResponseWriter, r *http.Request) {

	// ID
	id := presenters.GetRequestValue("id", r)

	//SET filter
	var qfilter model.QueryFilters

	qfilter.Filter = bson.M{"_id": bson.ObjectIdHex(id)}

	objectToReturn, errorFind := ui.usecase.GetOne(qfilter)
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
	var objectToCreate model.Species
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

	objectToCreate.Id = bson.NewObjectId()
	objectToCreate.Instance.Status = "ACTIVE"
	objectToCreate.Instance.CreatedAt = presenters.GetTimeNow()
	objectToCreate.Instance.ModifiedAt = presenters.GetTimeNow()

	_, errCreate := ui.usecase.Create(objectToCreate)
	if errCreate != nil {
		presenters.ReturnHttpError(errCreate, w, http.StatusInternalServerError)
		return
	}

	presenters.ReturnHttpPayload(bson.M{"success": true}, w)
}

func (ui *Interface) Delete(w http.ResponseWriter, r *http.Request) {
	//objectID
	id := presenters.GetRequestValue("id", r)

	//Set the filter to id
	queryFilters := model.QueryFilters{
		Filter: bson.M{"_id": bson.ObjectIdHex(id)},
	}
	errDelete := ui.usecase.Delete(queryFilters)
	if errDelete != nil {
		presenters.ReturnHttpError(errDelete, w, http.StatusInternalServerError)
		return
	}

	presenters.ReturnHttpPayload(bson.M{"succes": true}, w)
	return
}

func (ui *Interface) Update(w http.ResponseWriter, r *http.Request) {

	//Get Body update request
	var objectToUpdate model.Species
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		presenters.ReturnHttpError(err, w, http.StatusInternalServerError)
		return
	}
	errUnmarshal := json.Unmarshal(body, &objectToUpdate)
	if errUnmarshal != nil {
		presenters.ReturnHttpError(errUnmarshal, w, http.StatusInternalServerError)
		return
	}

	//Set the Filter to id
	queryFilters := model.QueryFilters{
		Filter: bson.M{"_id": objectToUpdate.Id},
	}

	errUpdated := ui.usecase.Update(&queryFilters, objectToUpdate)
	if errUpdated != nil {
		if strings.Contains(errUpdated.Error(), "E11000") {
			presenters.ReturnHttpError(errUpdated, w, http.StatusInternalServerError)
			return
		}
		presenters.ReturnHttpError(errUpdated, w, http.StatusInternalServerError)
		return
	}

	presenters.ReturnHttpPayload(bson.M{"success": true}, w)
	return
}
