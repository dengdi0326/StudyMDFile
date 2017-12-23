package main

import (
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"os"
)

func main() {
	const input  = "124 5678 1234567889"
	scanner := bufio.NewScanner(strings.NewReader(input))

	//自定义分割规则
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error){
		advance, token, err = bufio.ScanWords(data, atEOF)

		fmt.Println("token:", advance, token)

		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)

			fmt.Println(string(token))
			fmt.Println(strconv.ParseInt(string(token), 10, 32))
		}
		return
	}

	/*
	scanner.Split(bufio.ScanWords) //与上面一段效果相同
	*/

	scanner.Split(split)

	//判断是否读完
	for scanner.Scan(){
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Invalid input: %s", err)
	}

	const input2  = "the world"
	scanner = bufio.NewScanner(strings.NewReader(input2))

	scanner.Split(bufio.ScanWords)

	var cout int
	for scanner.Scan(){
		cout++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Inalid input: %s", err)
	}

	fmt.Print("the number of words: ", cout)

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	w.Flush()
}