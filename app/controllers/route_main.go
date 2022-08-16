package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "スレッド一覧", "layout", "top")
}

func thread(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "thread")
}

