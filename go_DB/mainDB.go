package main

import (
	"database/sql"
	"fmt"
	"go_docker/dao"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func run() error {

	dsn := mysqlDSN()
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		test(w, r, db)
	})
	http.ListenAndServe(":8080", r)

	return fmt.Errorf("dd")
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func test(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	ctx := r.Context()
	a, err := dao.Users(
		dao.UserWhere.ID.EQ("fff"),
	).One(ctx, db)
	if err != nil {
		w.Write([]byte("何もないおおおおおお"))
		return
	}
	w.Write([]byte(a.LastName))
	w.Write([]byte(a.FirstName))
}

// mysqlDSN MySQLの接続に必要なdata source nameを返す
func mysqlDSN() string {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)
}
