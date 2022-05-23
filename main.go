package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type utilisateur2 struct {
	prenom     string
	email      string
	CheckEmail string
	pays       string
	password   string
}

func utilisateur(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("enregistre.html"))
	pseudo := r.FormValue("pseudo")
	email := r.FormValue("Email")
	CheckEmail := r.FormValue("Email2")
	password := r.FormValue("password")
	nom := r.FormValue("Nom")
	prenom := r.FormValue("prenom")
	log.Print(pseudo + email + CheckEmail + password + nom + prenom)
	if r.FormValue("run") == "envoyer" {
		log.Print(CheckEmail)
		createUser(prenom, email, password, nom, pseudo)
	}
	tmpl.ExecuteTemplate(w, "menu", nil)
}
func createUser(firstname string, email string, password string, lastname string, pseudo string) {
	run, err := db.Query(fmt.Sprintf("INSERT INTO `users` (`Pseudo`,`Email`,`Password`,`Nom`,`Prenom`)  VALUES ('%s','%s','%s' ,'%s','%s')", pseudo, email, password, lastname, firstname))
	if err != nil {
		log.Println(err.Error())
	}
	defer run.Close()
}

func main() {
	directory := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", directory))
	directoryImg := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/", directoryImg))
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/cristiankayn")
	defer db.Close()
	insert, err := db.Query("INSERT INTO `agir` (`id_users`)  VALUES ('TEST' )")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("zdxfgghgh")
		tmpl := template.Must(template.ParseFiles("enregistre.html"))
		tmpl.Execute(rw, nil)
	})
	http.HandleFunc("/test", utilisateur)

	http.ListenAndServe("127.0.0.1:8888", nil)
}
