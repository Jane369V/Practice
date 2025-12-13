package auth

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "./auth.db")
	if err != nil {
		log.Fatal("Ошибка базы:", err)
	}

	// Таблица пользователей
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            surname TEXT,
            email TEXT UNIQUE,
            password TEXT,
            birth_date DATETIME,
            avatar_url TEXT
        );
    `)
	if err != nil {
		log.Fatal("Ошибка создания users:", err)
	}

	// Таблица refresh-token
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS refresh_token (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            token TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );
    `)
	if err != nil {
		log.Fatal("Ошибка создания refresh_token:", err)
	}

	log.Println("База данных подключена, таблицы готовы")

	// Создаём таблицу users
	createUsersTable := `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    surname TEXT,
    email TEXT UNIQUE,
    password TEXT,
    birth_date TEXT,
    avatar TEXT
);`

	_, err = DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Ошибка создания таблицы users:", err)
	}

}
