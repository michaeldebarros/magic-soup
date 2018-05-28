package db

import mgo "gopkg.in/mgo.v2"

//MgoSession Export (could have put this in another package, but the the connection is oly used by functions in controller package)
var MgoSession *mgo.Session

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	MgoSession = session

}
