package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	// Route
	http.HandleFunc("/", dailyBeautifulSentence)
	err := http.ListenAndServe(":23333", nil)
	if err != nil {
		log.Println(err)
	}
}

func dailyBeautifulSentence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/text")             //返回数据格式是text
	n := rand.Intn(150) + 1
	data := queryFromDB(n)
	if data == "" {
		io.WriteString(w, "nop!!!")
	} else {
		io.WriteString(w, data)
	}
}

func queryFromDB(n int) string {
	var s string
	db, err := sql.Open("mysql", "root:linglinger@tcp(127.0.0.1:3306)/goweb?charset=utf8")
	if err != nil {
		log.Println(err)
		db.Close()
		return ""
	}
	rows, err := db.Query("select sentence from davis where id=?", n)
	if err != nil {
		log.Println(err)
		db.Close()
		return ""
	}
	for rows.Next() {
		err := rows.Scan(&s)
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	err = rows.Close()
	if err != nil {
		log.Println(err)
	}
	db.Close()
	return s
}
