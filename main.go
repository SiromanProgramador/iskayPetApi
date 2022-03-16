package main

import (
	"challengeIskayPet/Boot"

	"challengeIskayPet/presenters"

	PetsRepo "challengeIskayPet/Domains/pets/entity/repository/iskayPet"
	PetsInterface "challengeIskayPet/Domains/pets/interface"
	PetsUsecase "challengeIskayPet/Domains/pets/usecase"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//initalized variables
var startTime time.Time
var db *mgo.Database

func main() {

	//Init DataBase
	db = MongoStart()
	router := loadRouter()
	Boot.Boot(db.Session)
	log.Fatal(http.ListenAndServe(":8000", router))

}

//start MongoDB
func MongoStart() *mgo.Database {
	session, err :=
		mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("challengeDB")

	return db

}

//charge router
func loadRouter() *mux.Router {

	router := mux.NewRouter()

	// Repos
	petsRepo := PetsRepo.NewIskayPetRepository(db.Session)

	// Usecases
	petsUsecase := PetsUsecase.NewUsecase(petsRepo)

	// Interfaces
	petsInterface := PetsInterface.PetsInterface(petsUsecase)

	//Pet Router
	router.HandleFunc("/creamascota", petsInterface.Create).Methods("POST")
	router.HandleFunc("/kpidemascotas", petsInterface.GetStatistics).Methods("GET")
	router.HandleFunc("/lismascotas", petsInterface.GetAll).Methods("GET")

	//Other Routes ================================
	router.HandleFunc("/", HealthCheckHandler)

	//Handle all OPTIONS Requests
	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			OptionsResponse(w, r)
		})

	return router
}

//Check the status and uptime of the server
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	lifeTime := bson.M{"started": startTime.Format("2006-01-02T15:04:05.999999-07:00"), "uptime": fmt.Sprint(math.Floor(time.Since(startTime).Seconds()*1000) / 1000)}
	presenters.ReturnHttpPayload(lifeTime, w)
}

//Empty Response for OPTIONS Request
func OptionsResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "limit, includes, skip, sort, filter, select, self, Accept, Accept-Language, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH, POST, GET, OPTIONS, PUT, DELETE")

	w.WriteHeader(200)
}
