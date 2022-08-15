package models

import (
	"log"
	"time"
)

type Thread struct {
	ID        int
	Title     string
	CreatedAt time.Time
	Contents  []Content
	Users     []User
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
