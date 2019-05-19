package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/BottleneckStudio/keepmotivat.in/app/controllers"
	"github.com/BottleneckStudio/keepmotivat.in/models"
	"github.com/BottleneckStudio/keepmotivat.in/server"
	tmpl "github.com/BottleneckStudio/keepmotivat.in/template"
	"github.com/go-chi/chi"
)

const (
	dbName = "keepmotivatin"
	dsn    = "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4"
)

func main() {

	router := chi.NewRouter()

	// Middleware here if ever.

	db, err := models.NewDatabase(dsn)
	if err != nil {
		// panic for now.
		panic(err.Error())
	}

	defer db.Close()

	db.Create(dbName)
	db.Use(dbName)
	db.Migrate()

	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		tpl := tmpl.New("./app/views/")

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tpl.Render(w, "hello.html", "HELLO WORLD ALL CAPS!"); err != nil {
			return
		}
	})

	router.Get("/", controllers.FeedController())
	router.Get("/tos", controllers.TermsOfServiceController())
	router.Get("/about", controllers.AboutController())
	router.Get("/privacy", controllers.PrivacyController())
	router.Get("/post/{postID}", controllers.PostController())

	s := server.New(":1333", router)
	go func() {
		s.Start()
	}()

	gracefulShutdown(s.HTTPServer)
}

func gracefulShutdown(srv *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Something went wrong with the graceful shutdown: %v", err)
	}
}
