package models

import (
	"log"
	"time"
)

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func GetCatetories() (categories []Category, err error) {
	cmd := `select id, name, created_at from categories`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var category Category
		err = rows.Scan(
			&category.ID,
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

func GetCategoryByCategoryID(id int) (category Category, err error) {
	cmd := `select id, name, created_at from categories where id = ?`

	err = Db.QueryRow(cmd, id).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt)

	if err != nil {
		log.Println(err)
	}

	return category, err
}

func GetCategory(id int) (category Category, err error) {
	cmd := `select id, name, created_at from categories where id = ?`

	err = Db.QueryRow(cmd, id).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt)

	return category, err
}

func GetCategoryName(threadID int) (name string, err error) {
	cmd := `select id, name, created_at from categories where id = ?`

	category := Category{}

	err = Db.QueryRow(cmd, threadID).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt)

	name = category.Name

	return name, err
}
