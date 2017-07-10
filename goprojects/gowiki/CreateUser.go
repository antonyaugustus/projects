package main

import
(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

const (
	DB_HOST = "tcp(192.168.50.4:3306)"
	DB_NAME = "social_network"
	DB_USER = "root"
	DB_PASS = "admin"
)

type User struct {
	ID int "json:id"
	Name  string "json:username"
	Email string "json:email"
	First string "json:first"
	Last  string "json:last"
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	NewUser := User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")

	output, err := json.Marshal(NewUser)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))
	if err != nil {
		fmt.Println("Error when connect database, the error is '%v'", err)
	}

	sql := "INSERT INTO users set user_nickname='" + NewUser.Name + "', user_first='" + NewUser.First + "', user_last='" + NewUser.Last + "', user_email='" + NewUser.Email + "'"
	q, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/user/create", CreateUser).Methods("GET")
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}