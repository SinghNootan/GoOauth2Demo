package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	
	oauthStateString = "random"
)
func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL : "http://localhost:8080/welcome",
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile", "email"},
		Endpoint: google.Endpoint,
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
							<body>	
								<a href="/login"> Google Log In </a>

							</body>
					</html>`
	fmt.Fprintf(w, htmlIndex)

}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)	
}

func handleGoogleWelcome(w http.ResponseWriter, r *http.Request) {
	content, err := getUserData(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "content: %s", content)
}
func getUserData(state, code string) ([]byte, error){
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code is not valid: %s", err.Error())
	}
	fmt.Println(token)
//	response, err := http.Get("https://www.googleapis.com/auth/userinfo?acces_token="+ token.AccessToken)
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %s", err.Error())
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body) 
	fmt.Println(string(data))
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return data, nil
}

func main() {

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/welcome", handleGoogleWelcome)

	http.ListenAndServe(":8080", nil)
}