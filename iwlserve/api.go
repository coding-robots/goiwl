package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type apiResponse struct {
	Writer     string `json:"writer"`
	Id         string `json:"id"`
	WriterLink string `json:"writer_link"`
	BadgeLink  string `json:"badge_link"`
	ShareLink  string `json:"share_link"`
}

type apiErrResponse struct {
	Error string `json:"error"`
}

func apiError(w http.ResponseWriter, errDesc string) {
	data, err := json.Marshal(&apiErrResponse{errDesc})
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Printf("api error: %s", err)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("api error: %s", err)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "expecting POST request", http.StatusBadRequest)
		return
	}
	// "client_id" is required but currently unused.
	if r.FormValue("client_id") == "" {
		apiError(w, "missing required argument 'client_id'")
		return
	}
	// "permalink" is required but currently unused.
	if r.FormValue("permalink") == "" {
		apiError(w, "missing required argument 'permalink'")
		return
	}
	// XXX "function" argument no longer supported.

	author, err := analyze(r.FormValue("text"))
	if err != nil {
		apiError(w, err.Error())
		return
	}
	resp := &apiResponse{
		Writer:     author,
		Id:         crcByAuthor[author],
		WriterLink: getAuthorURL(siteBaseURL, author),
		BadgeLink:  getBadgeURL(siteBaseURL, author),
		ShareLink:  getShareURL(siteBaseURL, author),
	}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Printf("api error: %s", err)
		return
	}
	// Write response.
	_, err = w.Write(data)
	if err != nil {
		log.Printf("api error: %s", err)
	}
}
