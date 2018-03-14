package main

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/create", CreateHandler)
	mux.HandleFunc("/get_list", GetListHandler)
	mux.HandleFunc("/update", UpdateHandler)
	mux.HandleFunc("/delete", DeleteHandler)

	http.ListenAndServe(":3000", mux)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var post Post
	r.ParseForm()
	for params, _ := range r.Form {
		if err := json.Unmarshal([]byte(params), &post); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	db, err := connectDb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db.Create(&post)

	w.Header().Add("Content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "http://127.0.0.1:8080")

	w.WriteHeader(http.StatusOK)
}

func GetListHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	db, err := connectDb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	postList := []Post{}
	db.Order("id desc").Find(&postList)

	content, err := json.Marshal(postList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "http://127.0.0.1:8080")

	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var post Post
	r.ParseForm()
	for params, _ := range r.Form {
		if err := json.Unmarshal([]byte(params), &post); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	db, err := connectDb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updateOption := Post{}
	updateOption.ID = post.ID
	db.First(&updateOption)
	db.Model(&updateOption).Update(&post)

	w.Header().Add("Content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "http://127.0.0.1:8080")

	w.WriteHeader(http.StatusOK)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var post Post
	r.ParseForm()
	for params, _ := range r.Form {
		if err := json.Unmarshal([]byte(params), &post); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	db, err := connectDb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	deleteOption := Post{}
	deleteOption.ID = post.ID
	db.First(&deleteOption)
	db.Delete(&deleteOption)

	w.Header().Add("Content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "http://127.0.0.1:8080")

	w.WriteHeader(http.StatusOK)
}

func connectDb() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_crud?loc=Asia%2FTokyo&parseTime=true&charset=utf8mb4")
	return db, err
}
