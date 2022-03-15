package Boot

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

//Here we are put all methos that we want to autorun when init the Api

func Boot(session *mgo.Session) {

	log.Println("[START] CHALLENGE ISKAYPET API")

}
