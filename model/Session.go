package model

import "gopkg.in/mgo.v2/bson"

//Session export
type Session struct {
	ID     bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserID string        `bson:"userID" json:"userID"`
}
