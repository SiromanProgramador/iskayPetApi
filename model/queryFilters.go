package model

import "gopkg.in/mgo.v2/bson"

type QueryFilters struct {
	Select bson.M   `json:"select,omitempty" bson:"select,omitempty"`
	Filter bson.M   `json:"filter,omitempty" bson:"filter,omitempty"`
	Sort   []string `json:"sort,omitempty" bson:"sort,omitempty"`
	Limit  int      `json:"limit,omitempty" bson:"limit,omitempty"`
	Skip   int      `json:"skip,omitempty" bson:"skip,omitempty"`
}
