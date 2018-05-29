package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Note struct {
	InternalId bson.ObjectId `bson:"_id"`
	Id         string        `json:"id"`
	Body       string        `json:"body"`
	CreatedOn  time.Time     `json:"createdOn"`
	UpdatedOn  time.Time     `json:"updatedOn"`
	UserId     string        `json:"userId"`
}
