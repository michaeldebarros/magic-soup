package controller

import (
	"fmt"
	"router/db"

	"gopkg.in/mgo.v2/bson"
)

//DeleteSoup export
func DeleteSoup(id []string, success chan bool) {
	session := db.MgoSession.Copy()
	defer session.Close()

	objID := bson.ObjectIdHex(id[0])
	c := session.DB("RECEPIES").C("soups")
	err := c.RemoveId(objID)
	if err != nil {
		fmt.Println(err)
	}
	success <- true
}
