package model

import (
	"challengeIskayPet/presenters"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const DBCOLLECTION_PETS = "pets"

type Pets struct {
	Pets []Pet `json:"pets,omitempty" bson:"pets,omitempty"`
}

type Pet struct {
	Id          bson.ObjectId       `json:"_id,omitempty," bson:"_id,omitempty"`
	Name        string              `json:"name,omitempty," bson:"name,omitempty"`
	Gender      string              `json:"gender,omitempty," bson:"gender,omitempty"`
	SpeciesId   bson.ObjectId       `json:"speciesId ,omitempty," bson:"speciesId,omitempty"`
	SpeciesInfo Species             `json:"speciesInfo ,omitempty," bson:"speciesInfo,omitempty"`
	Age         int                 `json:"age ,omitempty," bson:"age,omitempty"`
	DateOfBird  time.Time           `json:"dateOfBird ,omitempty," bson:"dateOfBird,omitempty"`
	Instance    presenters.Instance `json:"instance,omitempty," bson:"instance,omitempty"`
}
