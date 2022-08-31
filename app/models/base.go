package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"monta-channel/config"

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
	tableNameSession  = "sessions"
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
		title STRING,
		category_id int,
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
			created_at DATETIME)`, tableNameCategory)

	Db.Exec(cmdCate)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	Db.Exec(cmdS)

	// cmd := fmt.Sprintf(`insert into %s (name, created_at) values (?, ?)`, tableNameCategory)

	// Db.Exec(cmd, "漫画", time.Now())
	// Db.Exec(cmd, "スポーツ", time.Now())
	// Db.Exec(cmd, "健康・食事", time.Now())
	// Db.Exec(cmd, "社会・経済", time.Now())
	// Db.Exec(cmd, "エンターテイメント", time.Now())
	// Db.Exec(cmd, "お金", time.Now())
	// Db.Exec(cmd, "仕事", time.Now())
	// Db.Exec(cmd, "インターネット", time.Now())
	// Db.Exec(cmd, "趣味", time.Now())
	// Db.Exec(cmd, "恋愛・結婚", time.Now())
	// Db.Exec(cmd, "ファッション", time.Now())
	// Db.Exec(cmd, "生活", time.Now())
	// Db.Exec(cmd, "旅行・観光", time.Now())
	// Db.Exec(cmd, "その他", time.Now())
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
