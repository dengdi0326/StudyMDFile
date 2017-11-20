package main

import (
	"net/http"
	"strings"
	"github.com/stretchr/gomniauth"
	"fmt"
	"github.com/stretchr/objx"
	gomniauthcommon "github.com/stretchr/gomniauth/common"
	"crypto/md5"
	"io"
	"log"
)

type ChatUser interface {
	UniqueID() string
	AvatarURL()   string
}

type chatUser struct {
	gomniauthcommon.User
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServerHTTP(w http.ResponseWriter, r *http.Request){
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value== "" {

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)

	}else if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return

		panic(err.Error())
	}else {

	}
	h.next.ServeHTTP(w, r)
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next:handler}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]

	switch action {
	case "login":

		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get provider %s:%s", provider, err), http.StatusBadRequest)
			return
		}

		loginURL, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to GetBeginAuthURL for %s:%s ", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)

	case "callback":

		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get provider %s:%s", provider, err), http.StatusBadRequest)
			return
		}

		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to complete auth for %s:%S", provider, err), http.StatusInternalServerError)
			return
		}

		user, err := provider.GetUser(creds)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to complete auth for %s:%S", provider, err), http.StatusInternalServerError)
			return
		}
		chatUser := &chatUser{User: user}

		m := md5.New()
		io.WriteString(m, strings.ToLower(user.Email()))
		chatUser.uniqueID = fmt.Sprintf("%x", m.Sum(nil))

		avatarURL, err := avatars.GetAvatarURL(chatUser)
		if err != nil {
			log.Fatalln("Error when trying to GetAvatarURL", "-", err)
		}


		authCookieValue := objx.New(map[string]interface{}{
			"userid":ChatUser.UniqueID(),
			"name":user.Name(),
			"avatar_uri":avatarURL,
		}).MustBase64()

		http.SetCookie(w , &http.Cookie{
			Name:"auth",
			Value:authCookieValue,
			Path:"/",
		})

		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}