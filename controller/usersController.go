package usersController

import (
	"context"
	"fluke-hospital/db"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection = db.DbConnect().Database("fluke-hospital-db").Collection("users")

func CreateUser() *mongo.InsertOneResult {
	layoutDate := "2006-01-02"

	myBirthDate := "1999-05-11"
	convertBirthDate, _ := time.Parse(layoutDate, myBirthDate)

	user := bson.D{{"firstName", "Thanaphot"}, {"lastName", "Yasamuth"}, {"birthDate", convertBirthDate}}
	fmt.Println("user:", user)

	resultInsertUser, err := usersCollection.InsertOne(context.TODO(), user)

	if err != nil {
		panic(err)
	}

	return resultInsertUser
}

func ListUser() []primitive.M {
	filter := bson.D{}
	cursor, err := usersCollection.Find(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func FindUser() {
	objectId, err := primitive.ObjectIDFromHex("647da2dec7b05e5ee23fbf7a")
	if err != nil {
		fmt.Println("error:", err)
	}

	var result bson.M

	if err = usersCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("result:", result)
}
