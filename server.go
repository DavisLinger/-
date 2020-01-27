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
	n := rand.Int() % 150
	log.Println("request path : ", r.URL.Path)
	log.Println("local n is : ", n)
	db, err := sql.Open("mysql", "root:linglinger@tcp(127.0.0.1:3306)/goweb?charset=utf8")
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
		_, err = io.WriteString(w, s)
		if err != nil {
			log.Println("err:", err)
		}
	}
	err = rows.Close()
	if err != nil {
		log.Println(err)
	}
	db.Close()

}
