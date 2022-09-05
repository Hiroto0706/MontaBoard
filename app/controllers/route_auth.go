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

			popularThreads, err := models.GetPopularThreadsLimitFive()
			if err != nil {
				log.Println(err)
			}
			categories, err := models.GetCatetories()
			if err != nil {
				log.Println(err)
			}

			thread := models.Thread{}
			thread.Threads = threads
			thread.NewThreads = newThreads
			thread.PopularThreads = popularThreads
			thread.Categories = categories

			generateHTML(w, thread, "layout", "signup", "side-btn-if-not-login", "side-menu")
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
		popularThreads, err := models.GetPopularThreadsLimitFive()
		if err != nil {
			log.Println(err)
		}
		categories, err := models.GetCatetories()
		if err != nil {
			log.Println(err)
		}

		thread := models.Thread{}
		thread.Threads = threads
		thread.NewThreads = newThreads
		thread.PopularThreads = popularThreads
		thread.Categories = categories
		
		generateHTML(w, thread, "layout", "login", "side-btn-if-not-login", "side-menu")
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

func setting(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		threads, err := models.GetThreads()
		if err != nil {
			log.Println(err)
		}
		newThreads, err := models.GetNewThreadsLimitFive()
		if err != nil {
			log.Println(err)
		}

		loginUser, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		categories, err := models.GetCatetories()
		if err != nil {
			log.Println(err)
		}
		popularThreads, err := models.GetPopularThreadsLimitFive()
		if err != nil {
			log.Println(err)
		}

		thread := models.Thread{}
		thread.Threads = threads
		thread.NewThreads = newThreads
		thread.Categories = categories
		thread.PopularThreads = popularThreads
		thread.UserName = loginUser.Name
		thread.User = loginUser

		generateHTML(w, thread, "layout", "setting", "side-btn-if-login", "side-menu-if-login")
	}
}

func nameSetting(w http.ResponseWriter, r *http.Request, id int) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user, err := models.GetUser(id)
		if err != nil {
			log.Println(err)
		}

		user.Name = r.PostFormValue("name")
		// user.Email = r.PostFormValue("email")
		// user.Password = r.PostFormValue("password")

		// log.Println(user.Name)

		err = user.UpdateUserName()
		if err != nil {
			log.Println(err)
		}

		setting(w, r, id)
	}
}

func emailSetting(w http.ResponseWriter, r *http.Request, id int) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user, err := models.GetUser(id)
		if err != nil {
			log.Println(err)
		}

		// user.Name = r.PostFormValue("name")
		user.Email = r.PostFormValue("email")
		// user.Password = r.PostFormValue("password")

		// log.Println(user.Password)

		err = user.UpdateUserEmail()
		if err != nil {
			log.Println(err)
		}

		setting(w, r, id)
	}
}

func passwordSetting(w http.ResponseWriter, r *http.Request, id int) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user, err := models.GetUser(id)
		if err != nil {
			log.Println(err)
		}

		// user.Name = r.PostFormValue("name")
		// user.Email = r.PostFormValue("email")
		user.Password = r.PostFormValue("password")

		// log.Println(user.Password)

		err = user.UpdateUserPassword()
		if err != nil {
			log.Println(err)
		}

		// log.Println(user.Password)

		setting(w, r, id)
	}
}
