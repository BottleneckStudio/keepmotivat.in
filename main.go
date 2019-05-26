package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/BottleneckStudio/keepmotivat.in/app/controllers"
	"github.com/BottleneckStudio/keepmotivat.in/models"
	"github.com/BottleneckStudio/keepmotivat.in/server"
	tmpl "github.com/BottleneckStudio/keepmotivat.in/template"
	"github.com/go-chi/chi"
)

const (
	dbName  = "keepmotivatin"
	dsn     = "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4"
	certKey = "./certificates/localhost+1.pem"
	privKey = "./certificates/localhost+1-key.pem"
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

		if pusher, ok := w.(http.Pusher); ok {
			if err := pusher.Push("/assets/stylesheets/app.css", &http.PushOptions{
				Header: http.Header{"Content-Type": []string{"text/css"}},
			}); err != nil {
				log.Fatalf("Server push is not supported: %v", err)
			}
		}

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

	// get static files
	workDir, err := os.Getwd()
	if err != nil {
		log.Printf("Current Dir not found: %v", err)
		return
	}

	filesDir := filepath.Join(workDir, "app/data/assets")
	FileServer(router, "/assets", http.Dir(filesDir))

	s := server.New(":1333", router)
	go func() {
		// s.Start()
		s.StartTLS(certKey, privKey)
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

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
