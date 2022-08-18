package controllers

import (
	"log"
	"monta-channel/app/models"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			threads, err := models.GetThreads()
			if err != nil {
				log.Println(err)
			}
			newThreads, err := models.GetNewThreadsLimitFive()
			if err != nil {
				log.Println(err)
			}

			indexThreads := models.Thread{}
			indexThreads.Threads = threads
			indexThreads.NewThreads = newThreads

			generateHTML(w, indexThreads, "layout", "signup", "side-btn-if-not-login", "side-menu")
		} else {
			http.Redirect(w, r, "/index", 302)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}

		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		user, err = models.GetUserByEmail(user.Email)
		if err != nil {
			log.Println(err)
		}

		if user.Password == models.Encrypt(r.PostFormValue("password")) {
			session, err := user.CreateSession()
			if err != nil {
				log.Println(err)
			}

			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/", 302)
		} else {
			http.Redirect(w, r, "/signup", 302)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		threads, err := models.GetThreads()
		if err != nil {
			log.Println(err)
		}
		newThreads, err := models.GetNewThreadsLimitFive()
		if err != nil {
			log.Println(err)
		}

		indexThreads := models.Thread{}
		indexThreads.Threads = threads
		indexThreads.NewThreads = newThreads
		generateHTML(w, indexThreads, "layout", "login", "side-btn-if-not-login", "side-menu")
	} else {
		http.Redirect(w, r, "/index", 302)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}

	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}

	http.Redirect(w, r, "/", 302)
}
