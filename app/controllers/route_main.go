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
		thread.Categories = categories
		thread.PopularThreads = popularThreads

		generateHTML(w, thread, "layout", "top", "side-btn-if-not-login", "side-menu")
	} else {
		http.Redirect(w, r, "/index", 302)
	}
}

func thread(w http.ResponseWriter, r *http.Request, id int) {
	thread, err := models.GetThread(id)
	if err != nil {
		log.Println(err)
	}

	threads, err := models.GetThreads()
	if err != nil {
		log.Println(err)
	}

	contents, err := models.GetContentsByThreadID(id)
	if err != nil {
		log.Println(err)
	}

	for i, _ := range contents {
		contents[i].ContentIDInThread = i + 1
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

	for i, _ := range contents {
		user, _ := models.GetUser(contents[i].UserID)
		contents[i].UserName = user.Name
		contents[i].ContentIDInThread = i + 1
	}

	category, _ := models.GetCategory(thread.CategoryID)

	thread.Contents = contents
	thread.Threads = threads
	thread.NewThreads = newThreads
	thread.PopularThreads = popularThreads
	thread.Categories = categories
	thread.CategoryName = category.Name

	generateHTML(w, thread, "layout", "side-btn-if-not-login", "side-menu", "thread")
}

func category(w http.ResponseWriter, r *http.Request, id int) {
	thread := models.Thread{}

	threads, err := models.GetThreadsByCategory(id)
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

	category, _ := models.GetCategory(id)

	thread.Threads = threads
	thread.NewThreads = newThreads
	thread.PopularThreads = popularThreads
	thread.Categories = categories
	thread.CategoryName = category.Name

	generateHTML(w, thread, "layout", "side-btn-if-not-login", "side-menu", "category")
}

func index(w http.ResponseWriter, r *http.Request) {
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

		generateHTML(w, thread, "layout", "index", "side-btn-if-login", "side-menu-if-login")
	}
}

func indexThread(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		thread, _ := models.GetThread(id)

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
			user, _ := models.GetUser(contents[i].UserID)
			contents[i].UserName = user.Name
			contents[i].ContentIDInThread = i + 1
		}

		newThreads, err := models.GetNewThreadsLimitFive()
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

		category, _ := models.GetCategory(thread.CategoryID)

		thread.Contents = contents
		thread.Threads = threads
		thread.LoginUserID = loginUserId
		thread.NewThreads = newThreads
		thread.Categories = categories
		thread.PopularThreads = popularThreads
		thread.CategoryName = category.Name
		thread.UserName = loginUser.Name

		generateHTML(w, thread, "layout", "side-btn-if-login", "side-menu-if-login", "index-thread")
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

func contentUpdate(w http.ResponseWriter, r *http.Request, id int) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		content, err := models.GetContent(id)
		if err != nil {
			log.Println(err)
		}

		content.Content = r.PostFormValue("content")
		err = content.UpdateContent()
		if err != nil {
			log.Println(err)
		}

		indexThread(w, r, content.ThreadID)
	}
}

func contentDelete(w http.ResponseWriter, r *http.Request, id int) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		content, err := models.GetContent(id)
		if err != nil {
			log.Println(err)
		}
		err = content.DeleteContent()
		if err != nil {
			log.Println(err)
		}

		indexThread(w, r, content.ThreadID)
	}
}

func ThreadNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		categoryID, _ := strconv.Atoi(r.PostFormValue("categories"))

		category, err := models.GetCategoryByCategoryID(categoryID)
		if err != nil {
			log.Println(err)
		}

		err = models.CreateThread(r.PostFormValue("title"), category.ID)
		if err != nil {
			log.Println(err)
		}

		CategoryForm, _ := strconv.Atoi(r.PostFormValue("category-form"))

		if CategoryForm == 0 {
			http.Redirect(w, r, "/index", 302)
		} else if CategoryForm == 1 {
			indexCategory(w, r, categoryID)
		}
	}
}

func indexCategory(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		thread := models.Thread{}

		threads, err := models.GetThreadsByCategory(id)
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
			user, _ := models.GetUser(contents[i].UserID)
			contents[i].UserName = user.Name
			contents[i].ContentIDInThread = i + 1
		}

		newThreads, err := models.GetNewThreadsLimitFive()
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

		category, _ := models.GetCategory(id)

		thread.Contents = contents
		thread.Threads = threads
		thread.LoginUserID = loginUserId
		thread.NewThreads = newThreads
		thread.Categories = categories
		thread.PopularThreads = popularThreads
		thread.CategoryName = category.Name
		thread.UserName = loginUser.Name

		generateHTML(w, thread, "layout", "side-btn-if-login", "side-menu-if-login", "index-category")
	}
}

