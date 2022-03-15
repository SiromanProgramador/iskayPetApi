package model

import (
	"challengeIskayPet/presenters"

	"gopkg.in/mgo.v2/bson"
)

const DBCOLLECTION_SPECIES = "species"

type Species struct {
	Id       bson.ObjectId       `json:"_id,omitempty," bson:"_id,omitempty"`
	Name     string              `json:"name,omitempty," bson:"name,omitempty"`
	Instance presenters.Instance `json:"instance,omitempty," bson:"instance,omitempty"`
}
