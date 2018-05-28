package model

import "gopkg.in/mgo.v2/bson"

//User export
type User struct {
	ID            bson.ObjectId   `bson:"_id,omitempty" json:"_id"`
	Login         string          `bson:"login" json:"login"`
	Password      []byte          `bson:"password" json:"password"`
	FavoriteSoups []bson.ObjectId `bson:"favoriteSoups" json:"favoriteSoups"`
}
