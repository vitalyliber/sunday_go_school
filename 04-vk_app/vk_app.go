package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

const redirectUrl string = "http://localhost:8000/auth/callback"

func AuthCallback(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := r.FormValue("code")

	GetAccessToken(w, code)
}

func RedirectToVk(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vk_auth_link := fmt.Sprintf("https://oauth.vk.com/authorize?client_id=5833423&redirect_uri=%s&display=page", redirectUrl)
	http.Redirect(w, r, vk_auth_link, http.StatusSeeOther)
}

func GetAccessToken(w http.ResponseWriter, code string) {
	access_token_url := fmt.Sprintf("https://oauth.vk.com/access_token?client_id=5833423&client_secret=PU5YLvFySADmaWvGoljL&code=%s&redirect_uri=%s", code, redirectUrl)

	response, err := http.Get(access_token_url)

	if err != nil {
		fmt.Fprintf(w, "Error: %s! \n", err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Fprintf(w, "Error: %s! \n", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	access_token := data["access_token"].(string)

	fmt.Fprintf(w, "Access Token: %s! \n", access_token)
}

func main() {
	router := httprouter.New()
	router.GET("/auth/callback", AuthCallback)
	router.GET("/", RedirectToVk)

	log.Fatal(http.ListenAndServe(":8000", router))
}
