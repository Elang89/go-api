package common

import (
	"gopkg.in/mgo.v2"
)

type Database struct {
	*mgo.Database
}

var DB *mgo.Database

func Init() *mgo.Database {

}
