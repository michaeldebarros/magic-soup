package main

import (
	"fmt"
	"html/template"
	"log"
	"magicSoup/controller"
	"magicSoup/db"
	"magicSoup/model"
	"magicSoup/usersession"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var indexTmpl = template.Must(template.ParseFiles("./static/index.html"))
var loginTmpl = template.Must(template.ParseFiles("./static/login.html"))

func main() {
	defer db.MgoSession.Close()
	router := httprouter.New()
	router.HandleMethodNotAllowed = false //prevent router from sending 405 to request to same rout
	router.GET("/", indexHandler)
	router.GET("/login", loginGetHandler)
	router.POST("/login", loginPostHandler)
	router.GET("/logout", logOutHandler)
	router.POST("/newsoup", newSoupHandler)
	router.POST("/delete", deleteSoupHandler)
	router.GET("/assets/*filePath", staticHandler)
	log.Fatal(http.ListenAndServe(":8080", usersession.LoginWall(router)))
	//log.Fatal(http.ListenAndServe(":8080", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := controller.GetAllSoups()
	if err != nil {
		panic(err)
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		//handle with message to the user
	}

	u, ok := usersession.SessionMAP[cookie.Value]
	if !ok {
		//handle this with a message to the user
	}

	res := model.Response{
		UserID: u,
		Soups:  results,
	}
	indexTmpl.Execute(w, res)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loginTmpl.Execute(w, nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	message := make(chan string)
	cookiePointer := make(chan *http.Cookie)
	go controller.UserLogin(r.PostForm["login"], r.PostForm["password"], message, cookiePointer)

	//receive cookie from channel and put in variable
	cookieToSet := <-cookiePointer
	//set the cookie
	if cookieToSet == nil {
		messageToPrint := <-message
		res := model.Response{
			SuccessMessage: messageToPrint,
		}
		loginTmpl.Execute(w, res)
		return
	}

	http.SetCookie(w, cookieToSet)

	//receive messsage from channel
	//This message will be used for toast message/notifications
	//for now just print

	http.Redirect(w, r, "/", 302)
}

func logOutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookieToBeDeleted, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
	}

	//delete session from SessionMAP and database
	success := make(chan bool)

	go usersession.DeleteSession(cookieToBeDeleted.Value, success)

	//wait for OK

	ok := <-success

	var message string

	if ok == true {
		//send new empty cookie
		newCookie := http.Cookie{
			Name:   "session",
			Value:  cookieToBeDeleted.Value,
			MaxAge: -1,
		}
		http.SetCookie(w, &newCookie)
		http.Redirect(w, r, "/login", 302) //fix this redirect
	} else if ok == false {
		message = "Problem Logging Out. Try Again."
		res := model.Response{
			SuccessMessage: message,
		}
		indexTmpl.Execute(w, res)
	}
}

func newSoupHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Parse body
	r.ParseForm()

	//get cookie and pass
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
	}

	//create success channel
	successChan := make(chan bool)

	//call function that adds soup to db
	go controller.AddNewSoup(r.PostForm["name"], r.PostForm["origin"], r.PostForm["ingredients"], r.PostForm["spicy"], cookie.Value, successChan)

	//receive success bool message
	success := <-successChan

	//write to template
	if success == true {
		http.Redirect(w, r, "/", 302)
	} else {

		res := model.Response{
			SuccessMessage: "There was a problem creating the soup.",
		}
		indexTmpl.Execute(w, res)
	}

}

func deleteSoupHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	successChan := make(chan bool)
	go controller.DeleteSoup(r.PostForm["button"], successChan)

	success := <-successChan

	if success == true {
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/", 304)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staticFilePath := "./assets/" + ps.ByName("filePath")
	http.ServeFile(w, r, staticFilePath)
}
