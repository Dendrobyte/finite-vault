package db

import "go.mongodb.org/mongo-driver/bson/primitive"

/* Stubbing out schema and models within the Mongo database for general use and consistency */

type UserData struct {
	Name                 string  `bson:"name"`
	Email                string  `bson:"email"`
	Balance              float32 `bson:"balance,truncate"`
	DailyIncrement       float32 `bson:"daily_increment,truncate"`
	AutoIncrementBalance bool    `bson:"auto_increment_balance"`
	LastCheckin          int64   `bson:"last_checkin"`    // unix timestamp
	TzOffset             int64   `bson:"tz_offset"`       // also unix, offset from UTC
	TransactionIds       []primitive.ObjectID   `bson:"transaction_ids"` // TODO :)
}

type Transaction struct {
	Amount      float32 `bson:"amount"`
	Description string  `bson:"description"`
	CreationTimestamp  int64   `bson:"creation_ts"`
}
