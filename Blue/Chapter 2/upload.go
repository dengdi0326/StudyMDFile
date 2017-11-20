package main

import (
	"net/http"
	"io/ioutil"
	"path"
	"io"
)

func uploaderHandle(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userid")
	file, header, err := r.FormFile("avatarfile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := path.Join("avatar", userID+path.Ext(header.Filename))
	err := ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, "sucessful")
}
