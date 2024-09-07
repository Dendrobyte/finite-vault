package vault

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/Dendrobyte/finite_vault/db"
)

var (
	ErrGetBalance           = errors.New("failed to get user's balance")
	ErrDailyNumberIncr      = errors.New("error attempting to increment a user's daily balance")
	ErrNotEnoughTimeElapsed = errors.New("not enough time has elapsed to update the user's balance")
)

// TODO: Move this into the main Finite Vault go file for when this is all refactored out
//		 No point if there's only one function in here...? Thie files may need some re-org to avoid cyclic dependencies

func GetUserVaultBalance(w http.ResponseWriter, r *http.Request) {
	// TODO: Instead of passing email, extract email from the token!
	email := r.FormValue("email")
	userBal, err := db.GetUserBalance(email)
	if err != nil {
		log.Printf("Failed to get balance for user with email: %v\n", email)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: Verify this actually sends as an integer for the field...? It's JSON so it should be ok, but maybe marshaling from struct is generally better?
	respObj := []byte(fmt.Sprintf(`{"balance": %v}`, userBal)) // TODO: Public util function for one-off responses like this?
	json.NewEncoder(w).Encode(respObj)
}

// Takes a user's email and checks whether or not to increase their daily number based on "now"
func IncrementBalanceByDailyNumber(email string) (newBalance float32, err error) {
	// Pull the user data from the db
	userData, err := db.GetExistingUserData(email)
	if err != nil {
		log.Printf("%v - %v", ErrDailyNumberIncr, err)
		return 0.0, err
	}

	// TODO: Timezone offset calculation (TzOffset exists on user field, but that comes later)
	// Make sure enough time has elapsed since the last update'
	lastCheckinTime := time.Unix(userData.LastCheckin, 0)
	currTime := time.Now()
	timeElapsedHrs := currTime.Sub(lastCheckinTime).Hours()
	if timeElapsedHrs <= 24 {
		return userData.Balance, ErrNotEnoughTimeElapsed
	}

	// TODO: Toggle based on user setting
	// Get how many days have passed, e.g. no check-in for a day so update daily bal * 2
	daysSinceLastUpdate := float32(math.Floor(timeElapsedHrs / 24)) // Ignore remainder, we may still be risking an extra day

	// Update the balance if so, and by the approprate amount\
	incrAmount := daysSinceLastUpdate * userData.DailyIncrement
	newBalance, err = db.UpdateUserBalance(userData, incrAmount)
	if err != nil {
		log.Printf("%v - %v", ErrDailyNumberIncr, err)
		return 0.0, err
	}

	// Write the new timestamp back into the user object
	// TODO: Am I just paranoid or does this feel like so many read/writes? We pull all the data so I think we're ok
	//		 Maybe one optimization is to update balance and last checkin at the same time, that way we don't commit an update if the checkin write doesn't work
	//		 Yea... That should be done. Hence no error variable either right now.
	err = db.UpdateUserLastCheckin(userData.Email, currTime.Unix())
	if err != nil {
		return newBalance, errors.New("issue with writing user last checkin :/")
	}

	return newBalance, nil
}

// Retrieve all transactions for a usert
func GetAllUserTransactions(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	userData, err := db.GetExistingUserData(email) // TODO: For the third time- pull email from token!! It means we don't have to "validate" with a db call lol, this is unnecessary
	if err != nil {
		log.Printf("Get rid of me! Error with all user transactions\n")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tnxObjs, err := db.GetAllUserTransactions(userData)
	if err != nil {
		log.Printf("error retrieving transactions: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal to JSON for the frontend
	json.NewEncoder(w).Encode(tnxObjs)
}

// Log a new user transaction
func PostNewUserTransaction(w http.ResponseWriter, r *http.Request) {
	// Pull the user data from the db
	email := r.FormValue("email")
	amount := r.FormValue("tnx_amount")
	tnxAmount, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		log.Printf("Transaction amount of %v is invalid!\n", amount)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	description := r.FormValue("tnx_description")  // TODO: Sanitize description
	userData, err := db.GetExistingUserData(email) // TODO: Again, extract email from the token here
	if err != nil {
		log.Printf("Failed to post a transaction for user with email: %v\n", email)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.CreateNewTransaction(userData.Email, float32(tnxAmount), description)
	if err != nil {
		log.Printf("Failed to create a new transaction: %v", err) // TODO: Make a proper error?
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Returning a 200 should be fine for our purposes, frontend can act accordingly
}
