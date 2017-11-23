package main

import (
	"archive/zip"
	"os"
	"io"
	"log"
)

func main() {
	f1, err := os.Open("/Users/yusank/go/src/github.com/sapxry/chapter2/gopher.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	f2, err := os.Open("/Users/yusank/go/src/github.com/sapxry/chapter2/todo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	f3, err := os.Open("/Users/yusank/go/src/github.com/sapxry/chapter2/readme.txt")
	if err != nil {
		log.Fatal(err)
	}

	var filearr = []*os.File{f1,f2,f3}

	dest := "/Users/yusank/go/src/github.com/sapxry/ex2/readme.zip"
	d, _ := os.Create(dest)
	defer d.Close()

	w := zip.NewWriter(d)
	defer w.Close()

	//filearr := make([]*os.File,3)
	//
	//var files = []struct {
	//	Name, Body string
	//}{
	//	{"readme.txt", "This archive contains some text files."},
	//	{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},

	//}
	//for n, file := range files {
	//	f, err := os.Create(file.Name)
	//	if err != nil {
	//		log.Fatal("1",err)
	//	}
	//	_, err = f.Write([]byte(file.Body))
	//	if err != nil {
	//		log.Fatal("2",err)
	//	}
	//	filearr[n] = f
	//}

	for _, f := range filearr {
		err := compress(f, "", w)
		if err != nil {
			return
		}
	}
	return
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()//返回一个 描述文件的 fileinfo 类型值
	if err != nil {
		return err
	}

	//isDir 判断是否是一个目录
	if info.IsDir() {

		prefix := prefix + "/" + info.Name() // 加上前缀
		fileinfo, err := file.Readdir(-1)//返回目录中所有文件对象的fileinfo，以切片形式返回
		if err != nil {
			return err
		}

		for _, fi := range fileinfo {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	}else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

