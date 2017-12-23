package main

import (
	"bytes"
	"compress/gzip"
	"time"
	"log"
	"io"
	"os"
	"fmt"
)

func main() {
	var b bytes.Buffer
	zw := gzip.NewWriter(&b)

	file := []struct{
		Name     string
		Modtime  time.Time
		Data     string
	}{
		{"first.txt", time.Now(), "go package/compress/gzip--example one"},
		{"second.txt", time.Now(), "go package/compress/gzip--example two"},
	}

	for _, File := range file {
		zw.Name    = File.Name
		zw.ModTime = File.Modtime
		if _, err := zw.Write([]byte(File.Data)); err != nil {
			log.Fatal(err)
		}

		err := zw.Close()
		if err != nil {
			log.Fatal(err)
		}

		zw.Reset(&b)
	}

	zr, err := gzip.NewReader(&b)
	if err != nil {
		log.Fatal(err)
	}

	for {
		zr.Multistream(false)
		fmt.Println(zr.Name, zr.ModTime)

		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}

		fmt.Println()

		err = zr.Reset(&b)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	_ = zr.Reset(&b)
}