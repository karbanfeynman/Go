package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"sync"

	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	seed_URL string
	dup_URLs = map[string]bool{}
	// Protect access to dup
	mu       sync.Mutex
	producer = 3
	consumer = 1
	db       *sql.DB
)

func panicIF(err error) {
	if err != nil {
		panic(err)
	}
}

func setupDB() *sql.DB {
	db_sqlite3, err := sql.Open("sqlite3", "/home/karban/SQLite/user.db")
	panicIF(err)

	return db_sqlite3
}

/* error handler function */
func err_Handler(ctx *fetchbot.Context, res *http.Response, err error) {
	fmt.Printf("[ERR] %s %s - %s\n", ctx.Cmd.Method(), ctx.Cmd.URL(), err)
}

func get_Handler(ctx *fetchbot.Context, res *http.Response, err error) {
	// Process the body to find the links
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		fmt.Printf("[ERR] %s %s - %s\n", ctx.Cmd.Method(), ctx.Cmd.URL(), err)
		return
	}
	// Enqueue all links as HEAD requests
	fmt.Printf("Enqueue: %d\n", producer)
	enqueueLinks(ctx, doc)
}

func head_Handler(ctx *fetchbot.Context, res *http.Response, err error) {
	if _, err := ctx.Q.SendStringGet(ctx.Cmd.URL().String()); err != nil {
		fmt.Printf("[ERR] %s %s - %s\n", ctx.Cmd.Method(), ctx.Cmd.URL(), err)
	}
}

func main() {
	var sql_Cmd string

	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	mux := fetchbot.NewMux()

	fmt.Printf("Register a ErrorHandler\n")
	mux.HandleErrors(fetchbot.HandlerFunc(err_Handler))

	// 1. Handle GET requests for html responses
	// 2. to parse the body and
	// 3. enqueue all links as HEAD requests.
	fmt.Printf("Register a GetHandler\n")
	mux.Response().Method("GET").ContentType("text/html").Handler(fetchbot.HandlerFunc(get_Handler))

	// 1. Handle HEAD requests for html responses coming from the source host.
	// 2. We don't want to crawl links from other hosts.
	fmt.Printf("Register a HeadHandler\n")
	mux.Response().Method("HEAD").Host("golang.org").ContentType("text/html").Handler(fetchbot.HandlerFunc(head_Handler))

	// Create the Fetcher, handle the logging first, then dispatch to the Muxer
	fmt.Printf("Register a LogHandler\n")
	h := logHandler(mux)

	fmt.Printf("New a fetchbot\n")
	f := fetchbot.New(h)

	/* Start processing */
	fmt.Printf("Start the fetchbot\n")
	q := f.Start()

	// Enqueue the seed, which is the first entry in the dup map
	fmt.Printf("Enqueue the seed\n")
	db = setupDB()

	sql_Cmd = sql_Cmd_Select(consumer)
	rows, err_DB := db.Query(sql_Cmd)
	if err_DB != nil {
		fmt.Printf("[ERR]DB select fail\n")
		panicIF(err_DB)
	}

	for rows.Next() {
		rows.Scan(&seed_URL)
		dup_URLs[seed_URL] = true

		_, err := q.SendStringGet(seed_URL)
		if err != nil {
			fmt.Printf("[ERR] GET %s - %s\n", seed_URL, err)
		}
	}

	fmt.Printf("Start fetch Process\n")
	q.Block()
	fmt.Printf("End the process\n")
}

// logHandler prints the fetch information
// and dispatches the call to the wrapped Handler.
func logHandler(wrapped fetchbot.Handler) fetchbot.Handler {
	return fetchbot.HandlerFunc(
		func(ctx *fetchbot.Context, res *http.Response, err error) {
			if err == nil {
				fmt.Printf("[%d] %s %s - %s\n", res.StatusCode,
					ctx.Cmd.Method(),
					ctx.Cmd.URL(),
					res.Header.Get("Content-Type"))
			}
			wrapped.Handle(ctx, res, err)
		})
}

func sql_Cmd_Select(count int) string {
	return fmt.Sprintf("SELECT URL FROM urls")
}

func sql_Cmd_Insert(count int, url_String string) string {
	return fmt.Sprintf("insert into urls(ID, URL) values(%d, '%s')",
		count, url_String)
}

/* 1. parse the whold links in the content of webpage
   2. save URLs into dup_URLs
*/
func enqueueLinks(ctx *fetchbot.Context, doc *goquery.Document) {
	var sql_Cmd string

	mu.Lock()
	doc.Find("a[href]").Each(
		func(i int, s *goquery.Selection) {
			val, _ := s.Attr("href")
			// Resolve address
			u, err := ctx.Cmd.URL().Parse(val)
			if err != nil {
				fmt.Printf("[ERR] resolve URL %s - %s\n", val, err)
				return
			}
			if !dup_URLs[u.String()] {
				if _, err := ctx.Q.SendStringHead(u.String()); err != nil {
					fmt.Printf("[ERR] enqueue head %s - %s\n", u, err)
				} else {
					dup_URLs[u.String()] = true
					sql_Cmd = sql_Cmd_Insert(producer, u.String())
					producer++
					db.Exec(fmt.Sprintf(sql_Cmd))
				}
			} else {
				fmt.Printf("URL: %s exist\n", u.String())
			}
		})
	mu.Unlock()
}
