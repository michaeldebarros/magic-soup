package controller

import (
	"fmt"
	"net/http"
	"router/db"
	"router/model"
	"router/usersession"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2/bson"
)

//UserLogin export
func UserLogin(login []string, password []string, message chan string, cookiePointers chan *http.Cookie) {
	session := db.MgoSession.Copy()
	defer session.Close()
	c := session.DB("RECEPIES").C("users")

	userByLogin := model.User{}

	c.Find(bson.M{"login": login[0]}).One(&userByLogin)

	//if there is no user in the db with that login
	if len(userByLogin.ID) == 0 {
		//create new user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password[0]), 4)
		if err != nil {
			fmt.Println(err)
		}
		//insert new user in db
		userToBeInserted := model.User{
			Login:    login[0],
			Password: hashedPassword,
			ID:       bson.NewObjectId(),
		}
		if err := c.Insert(userToBeInserted); err != nil {
			fmt.Println(err)
		}

		//get recently created userID from db
		//call cookieMaker with userID string
		newCookiePointer := usersession.InitSession(userToBeInserted.ID.Hex())

		cookiePointers <- newCookiePointer
	} else {
		//if there already is a user in the db make login
		err := bcrypt.CompareHashAndPassword(userByLogin.Password, []byte(password[0]))
		if err != nil {
			cookiePointers <- nil
			message <- "Unable to login user."
		} else {
			usersession.DeleteOldSessions(userByLogin.ID.Hex())
			newCookiePointer := usersession.InitSession(userByLogin.ID.Hex())
			cookiePointers <- newCookiePointer
		}
	}
}
