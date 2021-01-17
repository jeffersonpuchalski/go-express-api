package models

import (
	"fmt"
	"time" // get an object type

	"github.com/InterloperStudios/ims-api/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User Ims user object oriented.
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Name      string             `bson:"name"`
	fullName  string             `bson:"full_name"`
	token     string             `bson:"token"`
}

var db = database.New("")

// ReturnAllUsers Return all user registered in mongo with given filters
func ReturnAllUsers(filter bson.M) ([]*User, error) {
	var users []*User

	col := db.GetCollection("ims")
	cursor, err := col.Find(db.GetContext(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(db.GetContext()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Printf("Error on Decoding the document -> %s", err)
			return nil, err
		}
	}
	return users, err
}
