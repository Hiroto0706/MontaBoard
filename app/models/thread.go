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

func GetContentsByThreadID(id int) (contents []Content, err error) {
	cmd := `select id, content, user_id, thread_id, created_at from contents where thread_id = ?`
	rows, err := Db.Query(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var content Content
		err = rows.Scan(
			&content.ID,
			&content.Content,
			&content.UserID,
			&content.ThreadID,
			&content.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}

		contents = append(contents, content)
	}
	rows.Close()

	return contents, err
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
