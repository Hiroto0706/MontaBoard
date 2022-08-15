package models

import (
	"log"
	"time"
)

type Content struct {
	ID        int
	Content   string
	UserID    int
	ThreadID  int
	CreatedAt time.Time
}

func (u *User) CreateContent(content string, threadID int) (err error) {
	cmd := `insert into contents (
		content,
		user_id,
		thread_id,
		created_at) values (?, ?, ?, ?)`

	log.Println(u.ID)

	_, err = Db.Exec(cmd, content, u.ID, threadID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetContent(id int) (content Content, err error) {
	content = Content{}
	cmd := `select id, content, user_id, thread_id, created_at from contents where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&content.ID,
		&content.Content,
		&content.UserID,
		&content.ThreadID,
		&content.CreatedAt)

	return content, err
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

func (c *Content) UpdateContent() (err error) {
	cmd := `update contents set content = ? where user_id = ?`
	_, err = Db.Exec(cmd, c.Content, c.UserID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (c *Content) DeleteContent() (err error) {
	cmd := `delete from contents where id = ?`
	_, err = Db.Exec(cmd, c.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (t *Thread) DeleteContentsByThreadID() (err error) {
	cmd := `delete from contents where thread_id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
