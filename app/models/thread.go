package models

import (
	"log"
	"sort"
	"strconv"
	"time"
)

type Thread struct {
	ID             int
	Title          string
	CategoryID     int
	CreatedAt      time.Time
	Threads        []Thread
	NewThreads     []Thread
	PopularThreads []Thread
	Contents       []Content
	Categories     []Category
	LoginUserID    int
}

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func CreateThread(title string) (err error) {
	cmd := `insert into threads (
		title,
		created_at) values (?, ?)`

	_, err = Db.Exec(cmd, title, time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetThread(id int) (thread Thread, err error) {
	thread = Thread{}
	cmd := `select id, title, created_at from threads where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&thread.ID,
		&thread.Title,
		&thread.CreatedAt)

	return thread, err
}

func GetThreads() (threads []Thread, err error) {
	cmd := `select id, title, created_at from threads`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var thread Thread
		err = rows.Scan(
			&thread.ID,
			&thread.Title,
			&thread.CreatedAt)

		if err != nil {
			log.Println(err)
		}

		threads = append(threads, thread)
	}
	rows.Close()

	return threads, err
}

func GetNewThreadsLimitFive() (threads []Thread, err error) {
	cmd := `select id, title, created_at from threads order by created_at desc limit 5`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var thread Thread
		err = rows.Scan(
			&thread.ID,
			&thread.Title,
			&thread.CreatedAt)

		if err != nil {
			log.Println(err)
		}

		threads = append(threads, thread)
	}
	rows.Close()

	return threads, err
}

type Entry struct {
	rank  string
	value int
}
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].rank > l[j].rank)
	} else {
		return (l[i].value > l[j].value)
	}
}

func GetPopularThreadsLimitFive() (popularThreads []Thread, err error) {
	cmd := `select id, title, created_at from threads`
	var threads []Thread

	rows, _ := Db.Query(cmd)
	// if err != nil {
	// 	log.Println(err)
	// }

	for rows.Next() {
		var thread Thread
		_ = rows.Scan(
			&thread.ID,
			&thread.Title,
			&thread.CreatedAt)

		// if err != nil {
		// 	log.Println(err)
		// }

		threads = append(threads, thread)
	}
	rows.Close()

	Ranking := map[string]int{}
	for i, thread := range threads {
		contents, _ := GetContentsByThreadID(thread.ID)
		// if err != nil {
		// 	log.Println(err)
		// }
		contentsCount := len(contents)
		i++
		i := strconv.Itoa(i)
		Ranking[i] = contentsCount
	}

	r := List{}
	for k, v := range Ranking {
		e := Entry{k, v}
		r = append(r, e)
	}

	sort.Sort(r)

	for _, value := range r {
		rank, _ := strconv.Atoi(value.rank)
		addThread, _ := GetThread(rank)
		if err != nil {
			log.Println(err)
		}
		popularThreads = append(popularThreads, addThread)
	}

	return popularThreads[:5], err
}

func (t *Thread) UpdataThread() (err error) {
	cmd := `update threads set title = ? where id = ?`
	_, _ = Db.Exec(cmd, t.Title, t.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (t *Thread) DeleteThread() (err error) {
	cmd := `delete from threads where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.DeleteContentsByThreadID()
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetCatetories() (categories []Category, err error) {
	cmd := `select distinct name, created_at from categories`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var category Category
		err = rows.Scan(
			&category.Name,
			&category.CreatedAt)

		if err != nil {
			log.Println(err)
		}

		categories = append(categories, category)
	}
	rows.Close()

	return categories, err
}
