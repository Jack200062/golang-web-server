package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func HashPassword(password []byte) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.Execute(w, nil)
}

func new_user(w http.ResponseWriter, r *http.Request) {
	user := User{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/web_server")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `users` (`login`,`password`) VALUES('%s', '%s')", user.Login, string(HashPassword([]byte(user.Password)))))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func sign_up(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/sign_up.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func logged_in(w http.ResponseWriter, r *http.Request) {
	user := User{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/web_server")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pass := db.QueryRow("SELECT  password FROM users WHERE login=?", user.Login)
	var db_password string
	err = pass.Scan(&db_password)

	hashed_db_password := []byte(db_password)
	if bcrypt.CompareHashAndPassword(hashed_db_password, []byte(user.Password)) != nil {
		fmt.Fprintf(w, "Incorrect credentials.. Try again")
		w.WriteHeader(http.StatusUnauthorized)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleConnection() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	http.HandleFunc("/", home_page)
	http.HandleFunc("/sign_up", sign_up)
	http.HandleFunc("/new_user", new_user)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logged_in", logged_in)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleConnection()
}
