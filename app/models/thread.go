package models

import (
	"log"
	"time"
)

type Thread struct {
	ID        int
	Title     string
	CreatedAt time.Time
	Threads   []Thread
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

func (t *Thread) UpdataThread() (err error) {
	cmd := `update threads set title = ? where id = ?`
	_, err = Db.Exec(cmd, t.Title, t.ID)
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
