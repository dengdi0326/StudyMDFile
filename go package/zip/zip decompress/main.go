package main

import (
	"archive/zip"
	"os"
	"io"
	"log"
)

func main() {
	dest := "/Users/yusank/go/src/github.com/sapxry/zip compress/txt"

	r,err := zip.OpenReader("/Users/yusank/go/src/github.com/sapxry/zip compress/readme.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, file := range r.File {

		rc, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}

		filename := dest + file.Name

		w, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer w.Close()

		_,err = io.Copy(w, rc)
		if err != nil {
			log.Fatal(err)
		}
	}
}

