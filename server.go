package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func main() {
	data := InitRead()
	InitWrite(data)
	log.Println("write ok!")
	//Route
	http.HandleFunc("/", dailyBeautifulSentence)
	err := http.ListenAndServe(":23333", nil)
	if err != nil {
		log.Println(err)
	}
}

func dailyBeautifulSentence(w http.ResponseWriter, r *http.Request) {
	n := rand.Int() % 150
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/goweb?charset=utf8")
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	rows, err := db.Query("select sentence from davis where id=?", n)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	for rows.Next() {
		var s string
		err := rows.Scan(&s)
		if err != nil {
			log.Println(err)
			return
		}
		io.WriteString(w, s)
	}
	rows.Close()
	db.Close()

}
func InitRead() []string {
	var sentence = make([]string, 0)
	//filePath := `H:\DavisLing\Documents\s1.txt`
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return nil
	}
	var filename string
	if runtime.GOOS == "windows" {
		filename = path + "\\s1.txt"
	} else {
		filename = path + "/s1.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		sentence = append(sentence, line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return nil
			}
		}
	}
	return sentence
}
func InitWrite(sentence []string) {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/goweb?charset=utf8")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	for _, val := range sentence {
		_, err := db.Exec("INSERT into davis(sentence) values(?) ", val)
		if err != nil {
			log.Println(err)
		}
	}
}
