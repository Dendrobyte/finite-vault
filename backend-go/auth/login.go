package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Dendrobyte/finite_vault/db"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

// TODO: Adding google auth credentials
var (
	SIMPLE_LOGIN_CLIENT_ID     string = ""
	SIMPLE_LOGIN_CLIENT_SECRET string = ""
	JWT_KEY                    []byte = []byte{}
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	SIMPLE_LOGIN_CLIENT_ID = os.Getenv("SIMPLE_LOGIN_CLIENT_ID")
	SIMPLE_LOGIN_CLIENT_SECRET = os.Getenv("SIMPLE_LOGIN_CLIENT_SECRET")
	JWT_KEY = []byte(os.Getenv("JWT_KEY"))

	fmt.Println("-+- Auth module finished loading -+-")
}

/*
 * Functions to handle login authentication, namely OAuth callback functions and generation of web tokens
 */

// Main entry point for the /login route
func LoginByService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	if service == "google" {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("No support for google login yet\n"))
	} else if service == "proton" {
		r.ParseForm()
		token := r.FormValue("token")
		redirect_uri := r.FormValue("redirect_uri")
		userLoginInfo := LoginProton(token, redirect_uri)

		// Write encoded JSON to the w object. Once google is implemented, this can be moved outside the if/else blocks
		json.NewEncoder(w).Encode(userLoginInfo)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("That service is not supported"))
	}
}

// NOTE: Confusing name with the schema imo
type UserInfo struct {
	Email     string  `json:"email"`
	AuthToken string  `json:"auth_token"`
	Username  string  `json:"username"`
	Balance   float32 `json:"balance"`
}

// Payload for JWT information
// Information from the token is used for database access for security reasons
// Token should be in every request sent by the client
type JWTPayload struct {
	Email string
}

// Locally create a JWT with the email encoded in the token
// TODO: Token expiry? Include that in claims and verify timestamp later
func createJWT(email string) (s string, err error) {
	var t *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   "2h", // TODO: Make this current time + 2h or something configured as a constant
		})
	s, err = t.SignedString(JWT_KEY)
	if err != nil {
		fmt.Printf("JWT creation encountered an issue: %v\n", err)
		return "", errors.New("could not properly sign the created JWT with a key")
	}
	return
}

// Given a token, verify its validity
// If valid, it will return true and the user's email as a string
// If not valid, it will return false
// Error will only not be nil if something goes wrong with processing, and will default to false
func VerifyJWT(tokenString string) (valid bool, email string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return JWT_KEY, nil
	})

	if err != nil {
		return
	}

	// TODO: Validate the "exp" claim to make sure token is not expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		valid = true
		if e, ok_string := claims["email"].(string); !ok_string {
			err = errors.New("the email in the token is not valid")
		} else {
			email = e
			valid = true
		}
	}

	return
}

// The data received from SimpleLogin that we care about unmarshaling
type ProtonData struct {
	Token    string `json:"access_token"`
	Expiry   int32  `json:"expires_in"`
	UserData struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"user"`
	Error string `json:"error"`
}

// Verify user login with Simple Login (Proton) and validate
func LoginProton(token string, redirect_uri string) UserInfo {
	log.Printf("Logging in user via proton with token %v\n", token)
	simpleLoginData := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {token},
		"redirect_uri":  {redirect_uri},
		"client_id":     {SIMPLE_LOGIN_CLIENT_ID},
		"client_secret": {SIMPLE_LOGIN_CLIENT_SECRET},
	}

	protonResp, err := http.PostForm("https://app.simplelogin.io/oauth2/token", simpleLoginData)
	if err != nil {
		fmt.Println("Issue with post request to SimpleLogin")
		return UserInfo{}
	}
	defer protonResp.Body.Close()

	data := ProtonData{}
	json.NewDecoder(protonResp.Body).Decode(&data)
	if data.Error != "" {
		// TODO: Handle this properly
		return UserInfo{Username: "Invalid! Errored out on backend"}
	}

	// TODO: Fetch initial balance from mongo as well, other function
	userData := db.GetUser(data.UserData.Email, data.UserData.Name)
	jwt, err := createJWT(data.UserData.Email)
	if err != nil {
		return UserInfo{} // TODO: Properly bubble up errors here
	}

	log.Printf("Logged in user %v (name: %v)", data.UserData.Email, data.UserData.Name)
	return UserInfo{userData.Email, jwt, userData.Name, userData.Balance}
}
