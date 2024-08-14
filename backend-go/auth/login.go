package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

// TODO: Adding google auth credentials
var (
	SIMPLE_LOGIN_CLIENT_ID     string = ""
	SIMPLE_LOGIN_CLIENT_SECRET string = ""
)

func init() {
	path, _ := os.Getwd() // Since it's not in the same place as the executable, get path to working directory
	fmt.Println("The dir: ", path)
	if err := godotenv.Load(path + "/auth/.env"); err != nil {
		fmt.Println(err)

	}

	SIMPLE_LOGIN_CLIENT_ID = os.Getenv("SIMPLE_LOGIN_CLIENT_ID")
	SIMPLE_LOGIN_CLIENT_SECRET = os.Getenv("SIMPLE_LOGIN_CLIENT_SECRET")

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
		json.NewEncoder(w).Encode(userLoginInfo)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("That service is not supported"))
	}
}

type UserInfo struct {
	AuthToken string  `json:"auth_token"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Balance   float32 `json:"balance"`
}

// Payload for JWT information
// Information from the token is used for database access for security reasons
// Token should be in every request sent by the client
type JWTPayload struct {
	Email string
}

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
	fmt.Println(token)
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
	// TODO: Logging... :I

	// TODO: Incorporate JWT stuff for authtoken
	// TODO: Fetch initial balance from mongo as well, other function
	return UserInfo{data.Token, data.UserData.Name, data.UserData.Email, 0}
}
