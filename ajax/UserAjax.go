package ajax

import (
	"GoOnlineJudge/classes"
	"GoOnlineJudge/config"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type result struct {
	Uid       string
	Ok        int
	Privilege int
	Status    int
}

type UserAjax struct {
}

func (this *UserAjax) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("User Login")
	w.Header().Set("content-type", "application/json")

	r.ParseForm()

	response, _ := http.PostForm(config.Host+"/user/login", r.PostForm)
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		jsonBody := &result{}
		json.Unmarshal(body, jsonBody)

		if jsonBody.Ok == 1 {
			s := &classes.Session{Name: "uid", Value: jsonBody.Uid}
			s.Set(w, r)
		}

		w.Write(body)
	} else {
		log.Println("User Login Response Error")
	}
}

func (this *UserAjax) Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("User Logout")
	w.Header().Set("content-type", "application/json")

	r.ParseForm()

	response, _ := http.PostForm(config.Host+"/user/logout", r.PostForm)
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		jsonBody := &result{}
		json.Unmarshal(body, jsonBody)

		if jsonBody.Ok == 1 {
			s := &classes.Session{Name: "uid", Value: jsonBody.Uid}
			s.Set(w, r)
		}

		w.Write(body)
	} else {
		log.Println("User Logout Response Error")
	}
}
