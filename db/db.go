package db

import mgo "gopkg.in/mgo.v2"

//MgoSession Export
var MgoSession *mgo.Session

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	MgoSession = session
}
