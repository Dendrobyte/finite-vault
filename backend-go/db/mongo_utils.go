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

// Given an email and name, retrieves a user from the database
// If no document is found, we create the user
func GetUserDataOrCreate(email string, name string) (user UserData) {
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

// Given a user email, get the user information from the database
func GetExistingUserData(email string) (user UserData, err error) {
	userColl := mongoClient.Database(database).Collection("users") // TODO: Can this move out...?

	err = userColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return UserData{}, fmt.Errorf("error retrieving existing user from mongo: %v", err)
	}

	return user, nil
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

// Retrieves a user's balance from mongo, filtered by email
func GetUserBalance(email string) (float32, error) {
	var result UserData

	userColl := mongoClient.Database(database).Collection("users")
	opts := options.FindOne().SetProjection(bson.D{{Key: "balance", Value: 1}})
	err := userColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}, opts).Decode(result)
	if err != nil {
		return 0.00, fmt.Errorf("getting user balance failed on mongo: %v", err)
	}

	return result.Balance, nil
}

// Updates a user's balance and returns the updated balance
func UpdateUserBalance(data UserData, change float32) (float32, error) {

	userColl := mongoClient.Database(database).Collection("users")
	newBalance := data.Balance - change
	filter := bson.D{{Key: "email", Value: data.Email}}
	update := bson.D{{Key: "balance", Value: newBalance}}
	_, err := userColl.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return data.Balance, err
	}

	return newBalance, nil
}

// Retrieves a user's daily increment number from mongo, filtered by email
// TODO: May not be needed, only use just pulls entire user anyway
func GetUserDailyIncrement(email string) (float32, error) {
	var result UserData

	userColl := mongoClient.Database(database).Collection("users")
	opts := options.FindOne().SetProjection(bson.D{{Key: "balance", Value: 1}})
	err := userColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}, opts).Decode(result)
	if err != nil {
		return 0.00, fmt.Errorf("getting user balance failed on mongo: %v", err)
	}

	return result.Balance, nil
}

func UpdateUserLastCheckin(email string, timestamp int64) (err error) {

	userColl := mongoClient.Database(database).Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "last_checkin", Value: timestamp}}
	_, err = userColl.UpdateOne(context.TODO(), filter, update)

	return
}
