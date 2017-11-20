package main

import (
	"errors"
	"io/ioutil"
	"strings"
	"path/filepath"
)

var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

type Avatar interface {
	GetAvatarURL(c ChatUser) (string, error)
}

type tryAvatar []Avatar
func(a tryAvatar) GetAvatarURL (u chatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil{
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}
// method one
type AuthAvatar struct {}
var UserAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}
	return u.AvatarURL(), nil
}

//method two
type GravatarAvatar struct {}
var UserGravatarAvatar GravatarAvatar

func(GravatarAvatar) GetAvatarURL (u ChatUser) (string, error) {
	return "www.gravatar.com/avatar/" + u.UniqueID(), nil
}

//method three
type FileSystemAvatar struct {}
var UserFileSystemAvatar FileSystemAvatar

func(FileSystemAvatar) GetAvatarURL (u ChatUser)(string, error) {
	files, err := ioutil.ReadDir("avatars")
	if err != nil {
		return "", ErrNoAvatarURL
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fname := file.Name()
		if u.UniqueID() == strings.TrimSuffix(fname, filepath.Ext(fname)){
			return "avatar/" + file.Name(), nil
		}
	}
	return "", ErrNoAvatarURL
}