package usersession

import (
	"net/http"
	"router/db"
	"router/model"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//LoginWall export
func LoginWall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check all routes besides login and assets if cookie is present
		if r.URL.String() != "/login" && !strings.HasPrefix(r.URL.String(), "/assets") {
			cookie, err := r.Cookie("session")
			if err != nil {
				http.Redirect(w, r, "/login", 302)
				return
			}
			_, ok := SessionMAP[cookie.Value]
			if !ok {
				//check db for session in db
				session := db.MgoSession.Copy()
				defer session.Close()

				//look for session in db
				restoredSession := model.Session{}
				c := session.DB("RECEPIES").C("sessions")
				c.FindId(bson.ObjectIdHex(cookie.Value)).One(&restoredSession)

				//check if the struct is empty, meaning no session found in db
				if (model.Session{}) == restoredSession {
					http.Redirect(w, r, "/login", 302)
				}

				//found session on db, write to SessionMap
				SessionMAP[cookie.Value] = restoredSession.UserID
			}
		}
		next.ServeHTTP(w, r)
	})
}
