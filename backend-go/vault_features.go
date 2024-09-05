package main

import "github.com/Dendrobyte/finite_vault/db"

// TODO: Move this into the main Finite Vault go file for when this is all refactored out

// Takes a user's email and checks whether or not to increase their daily number based on "now"
func IncrementBalanceByDailyNumber(data db.UserData) {
	// TODO: Throw the daily number on the user data object to avoid "extra" db call? I guess we just have to send the whole user object to the backend...
	//	     but I don't want to hold on to copies of all users here... Start from the other direction I guess

}
