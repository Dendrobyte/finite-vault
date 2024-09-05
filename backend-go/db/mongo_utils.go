package db

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_URI   string
	mongoClient *mongo.Client
	database    string = "test"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	MONGO_URI = os.Getenv("MONGO_URI")

	fmt.Println("-+- Mongo utils finished loading -+-")
}

// Initialize the mongo db connection
// Is called in the main function, should not be called otherwise
func InitMongoDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URI))

	if err != nil {
		panic(err)
	}

	mongoClient = client
}

// Given an email, retrieves a user from the database
// Name is not used for database calls, but for creation in case none provided
func GetUser(email string, name string) (user UserData) {
	userColl := mongoClient.Database(database).Collection("users")

	err := userColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		user, _ = createUser(email, name, userColl)
		return
	} else {
		if err != nil {
			fmt.Printf("There was an error finding user with email %s: %v", email, err)
		}
	}

	return user
}

// Creates a new user, which happens if an email is not found in the database
func createUser(email string, name string, userColl *mongo.Collection) (newUserData UserData, err error) {
	newUserData = UserData{Email: email, Name: name, Balance: 0.00}
	_, err = userColl.InsertOne(context.TODO(), newUserData)
	if err != nil {
		fmt.Printf("Error creating a new user for email %s: %v", email, err)
		return
	}

	// It's been inserted into the database, so we can just return the information we would effectively retrieve normally (i.e. defaults)
	fmt.Printf("Created new user for email %v\n", email)
	return newUserData, nil
}

// Updates a user's balance and returns the updated balance
func UpdateUserBalance(data UserData, change float32) (float32, error) {

	userColl := mongoClient.Database(database).Collection("users")
	newBalance := data.Balance - change
	// TODO: Get their daily number, then increment the balance, then write it back. All in this function.
	filter := bson.D{{Key: "email", Value: data.Email}}
	update := bson.D{{Key: "balance", Value: newBalance}}
	_, err := userColl.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return data.Balance, err
	}

	return newBalance, nil
}
