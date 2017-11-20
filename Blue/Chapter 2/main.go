package main

import (
	"sync"
	"html/template"
	"net/http"
	"path/filepath"
	"github.com/stretchr/objx"
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"log"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func(t *templateHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func(){
		t.templ = template.Must(template.ParseFiles(filepath.Join("template", t.filename)))
	})

	date := map[string]interface{}{
		"host": r.Host,
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		date["UserDate"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, date)
}

var host = flag.String("host", ":8080", "The host of the application.")

var avatars Avatar = tryAvatar {
	UserAuthAvatar,
	UserFileSystemAvatar,
	UserGravatarAvatar,
}

func main() {
	flag.Parse()

	gomniauth.SetSecurityKey("asdfasdf")
	gomniauth.WithProviders(
		github.New("", "", ""),
	)

	r := newRoom(UserAuthAvatar)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request){
		http.SetCookie(w, &http.Cookie{
			Name:"auth",
			Value:"",
			Path:"/",
			MaxAge:-1,
		})

		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandle)

	http.Handle("/avatars/",
		http.StripPrefix("/avatars/",
			http.FileServer(http.Dir("./avatars"))))

	go r.run()

	log.Println("Starting web server on", *host)
	if err := http.ListenAndServe(*host, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

