package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"runtime"
	"os"
	"os/exec"
)

type Users struct {
	Id   int
	Name string
}

type Image struct {
	Name string
}

func handler_Root(r render.Render) {
	r.HTML(200, "index", nil)
}

func capture_image() {
	cmd := exec.Command("gphoto2", "--capture-image")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func handler(r render.Render, params martini.Params) {
	switch params["name"] {
	case "index":
		images := []Image{}
		files, _ := ioutil.ReadDir("/home/karban/Go/Go/templates/html5/IMAGE")
		for _, f := range files {
			i := Image{}
			i.Name = f.Name()
			images = append(images, i)
			fmt.Printf("[Karban]name: %s\n", i.Name)
		}
		r.HTML(200, params["name"], images)
	case "capture":
		capture_image()
		r.HTML(200, "index", nil)
	default:
		r.HTML(200, params["name"], nil)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := martini.Classic()

	/* set host and port */
	m.Use(render.Renderer(render.Options{
		Directory: "templates/html5",
		Layout:    "layout",
	}))
	m.Use(martini.Static("templates/html5/"))

	m.Get("/", handler_Root)
	m.Get("/:name", handler)

	m.Run()
}
