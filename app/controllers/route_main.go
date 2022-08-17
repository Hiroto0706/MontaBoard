package controllers

import (
	"log"
	"monta-channel/app/models"
	"net/http"
	"strconv"
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

		generateHTML(w, indexThreads, "layout", "top", "side-btn-if-not-login", "side-menu")
	} else {
		http.Redirect(w, r, "/index", 302)
	}
}

func thread(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "thread", "side-btn-if-not-login", "side-menu")
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
		// newThreads, err := models.GetNewThreadsLimitFive()
		if err != nil {
			log.Println(err)
		}
		indexThreads := models.Thread{}
		indexThreads.Threads = threads
		// indexThreads.NewThreads = newThreads

		generateHTML(w, indexThreads, "layout", "index", "side-btn-if-login", "side-menu")
	}
}

func indexThread(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		thread, err := models.GetThread(id)
		if err != nil {
			log.Println(err)
		}

		threads, err := models.GetThreads()
		if err != nil {
			log.Println(err)
		}

		loginUser, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		loginUserId := loginUser.ID

		contents, err := models.GetContentsByThreadID(id)
		if err != nil {
			log.Println(err)
		}

		for i, _ := range contents {
			contents[i].LoginUserID = loginUserId
			contents[i].ContentIDInThread = i + 1
		}

		thread.Contents = contents
		thread.Threads = threads
		thread.LoginUserID = loginUserId

		generateHTML(w, thread, "layout", "side-btn-if-login", "side-menu", "index-thread")
	}
}

func contentNew(w http.ResponseWriter, r *http.Request, userId int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		contentThreadID, err := strconv.Atoi(r.PostFormValue("id"))

		content := models.Content{
			Content:  r.PostFormValue("content"),
			ThreadID: contentThreadID,
		}

		err = user.CreateContent(content.Content, content.ThreadID)
		if err != nil {
			log.Println(err)
		}

		indexThread(w, r, content.ThreadID)
	}
}

// func contentUpdate(w http.ResponseWriter, r *http.Request, id int){
// 	sess, err := session(w, r)
// 	if err != nil {
// 		http.Redirect(w, r, "/login", 302)
// 	}else {
// 		_, err := sess.GetUserBySession()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		content, err := models.GetContent(id)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// }
