package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"monta-channel/config"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser     = "users"
	tableNameContent  = "contents"
	tableNameThread   = "threads"
	tableNameCategory = "categories"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		Title STRING,
		created_at DATETIME)`, tableNameThread)

	Db.Exec(cmdT)

	cmdC := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		thread_id INTEGER,
		created_at DATETIME)`, tableNameContent)

	Db.Exec(cmdC)

	cmdCate := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name STRING,
			thread_id INTEGER,
			created_at DATETIME)`, tableNameCategory)

	Db.Exec(cmdCate)

	cmd := `insert into users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)`

	Db.Exec(cmd, createUUID(), "ジョルノジョバァーナ", "test@test.com", Encrypt("password"), time.Now())

	cmd = `insert into threads (title, created_at) values (?, ?)`

	Db.Exec(cmd, "スレッド1", time.Now())

	cmd = `insert into contents (content, user_id, thread_id, created_at) values (?, ?, ?, ?)`

	Db.Exec(cmd, "this is a test", 1, 1, time.Now())

	cmd = `insert into categories (name, thread_id, created_at) values (?, ?, ?)`

	Db.Exec(cmd, "カテゴリー1", 1, time.Now())
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
