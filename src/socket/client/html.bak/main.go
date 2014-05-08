package main

import (
	"html/template"
	"net/http"
	"time"

	"engine/conf"
	"engine/errors"
	"engine/log"
)

var (
	filenames = []string{"index.html"}
	// TODO:必须先执行Funcs再Parse
	t          = template.Must(template.New("socket.io").Funcs(template.FuncMap{"format": Format}).ParseFiles(filenames...))
	user       = New()
	publicPath = conf.RootPath() + "/public"
)

func Format(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// index
func Index(w http.ResponseWriter, r *http.Request) {
	// FindAll
	users := user.FindAll()

	if err := t.ExecuteTemplate(w, "index.html", users); err != nil {
		log.Warn(errors.As(err))
	}

	return
}

// index
func Add(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if len(name) == 0 {
		name = "Empty"
	}

	user.Add(name)

	users := user.FindAll()

	if err := t.ExecuteTemplate(w, "index.html", users); err != nil {
		log.Warn(errors.As(err))
	}

	return
}

// public
func Public(w http.ResponseWriter, r *http.Request) {
	handle := http.FileServer(http.Dir(publicPath))
	http.StripPrefix("/public", handle).ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", Public)
	http.HandleFunc("/index", Index)
	http.HandleFunc("/add", Add)

	// listen
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Warn(errors.As(err))
	}
}
