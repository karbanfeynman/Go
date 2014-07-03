package main

import (
	"database/sql"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
	"runtime"
)

type Users struct {
	Id   int
	Name string
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func SetupDB() *sql.DB {
	db, err := sql.Open("sqlite3", "/home/karban/SQLite/user.db")
	PanicIf(err)

	return db
}

func handler_Root(r render.Render, db *sql.DB) {
	rows, err := db.Query("SELECT ID, URL FROM href")
	PanicIf(err)
	defer rows.Close()

	users := []Users{}
	for rows.Next() {
		u := Users{}
		err := rows.Scan(&u.Id, &u.Name)
		PanicIf(err)
		users = append(users, u)
	}

	r.HTML(200, "index", nil)

}

/*
func handler(r render.Render, args martini.Params) {
}
*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := martini.Classic()
	m.Map(SetupDB())

	/* set host and port */
	m.Use(render.Renderer(render.Options{
		Directory: "templates/html5",
		Layout:    "layout",
	}))
	m.Use(martini.Static("templates/html5/"))

	m.Get("/", handler_Root)
	//	m.Get("/:id", handler)

	m.Run()
}
