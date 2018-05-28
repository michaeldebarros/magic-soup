package usersession

import (
	"fmt"
	"net/http"
	"router/db"
	"router/model"

	"gopkg.in/mgo.v2/bson"
)

//create the session MAP

//SessionMAP export
var SessionMAP map[string]string

func init() {
	m := make(map[string]string)
	SessionMAP = m
}

//InitSession export function
func InitSession(userIDHex string) *http.Cookie {
	//creat a session and insert in db
	session := db.MgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("sessions")

	//verify if there is a session in the db for that user and delete

	s := model.Session{
		ID:     bson.NewObjectId(),
		UserID: userIDHex,
	}

	if err := c.Insert(s); err != nil {
		fmt.Println(err)
	}

	sessionIDHex := s.ID.Hex()

	//write session to SessionMAP
	SessionMAP[sessionIDHex] = userIDHex

	//create cookie and return it
	newCookie := http.Cookie{
		Name:   "session",
		Value:  sessionIDHex,
		MaxAge: 1500,
	}

	return &newCookie
}

//DeleteSession export
func DeleteSession(sessionIDString string, success chan bool) {

	session := db.MgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("sessions")

	if err := c.RemoveId(bson.ObjectIdHex(sessionIDString)); err != nil {
		fmt.Println(err)
	}

	//delete from SessionMAP
	delete(SessionMAP, sessionIDString)

	success <- true

}

//DeleteOldSessions export
func DeleteOldSessions(userIDHex string) {
	session := db.MgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("sessions")

	//remove all sessions from certain user from db
	_, err := c.RemoveAll(bson.M{"userID": userIDHex})
	if err != nil {
		fmt.Println(err)
	}
}
