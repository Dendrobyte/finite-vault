package db

/* Stubbing out schema and models within the Mongo database for general use and consistency */

type UserData struct {
	Name                 string  `bson:"name"`
	Email                string  `bson:"email"`
	Balance              float32 `bson:"balance"`
	DailyIncrement       float32 `bson:"daily_increment"`
	AutoIncrementBalance bool    `bson:"auto_increment_balance"`
	LastCheckin          int64   `bson:"last_checkin"`    // unix timestamp
	TzOffset             int64   `bson:"tz_offset"`       // also unix, offset from UTC
	TransactionIds       []int   `bson:"transaction_ids"` // TODO :)
}
