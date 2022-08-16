package controllers

import (
	"log"
	"monta-channel/app/models"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {

		// threadsを取ってきている
		threads, err := models.GetThreads()
		if err != nil {
			log.Println(err)
		}
		indexThreads := models.Thread{}
		indexThreads.Threads = threads

		generateHTML(w, indexThreads, "layout", "top", "side-btn-if-not-login")
	} else {
		http.Redirect(w, r, "/index", 302)
	}
}

func thread(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "thread", "side-btn-if-not-login")
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		threads, err := models.GetThreads()
		if err != nil {
			log.Println(err)
		}
		indexThreads := models.Thread{}
		indexThreads.Threads = threads

		generateHTML(w, indexThreads, "layout", "index", "side-btn-if-login")
	}
}

func indexThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		generateHTML(w, nil, "layout", "side-btn-if-login", "index-thread")
	}
}
