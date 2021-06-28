package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	h "./src"
	"os"
	//_ "github.com/mattn/go-sqlite3"
)


func main() {

const (
	ConstLogFilename       string        = "test.txt"
)


	logFile, err := os.OpenFile(ConstLogFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("[ERROR] - Error opening file: %v", err)
		os.Exit(6)    //ErrorOpenFile
	}
	defer logFile.Close()
	h.Logger = log.New(logFile, "ERROR - ", log.LstdFlags)

	//db, err := sql.Open("sqlite3", "./")
	//Load the pages
	http.Handle("/", http.NotFoundHandler())
	http.HandleFunc("/home", h.Home)
	http.Handle("/home/", http.NotFoundHandler())
	http.HandleFunc("/homeLogged", h.HomeLogged)
	http.Handle("/homeLogged/", http.NotFoundHandler())
	http.HandleFunc("/login", h.Login)
	http.Handle("login/", http.NotFoundHandler())
	http.HandleFunc("/register", h.Register)
	http.Handle("/register/", http.NotFoundHandler())
	http.HandleFunc("/liked", h.Liked)
	http.Handle("/liked/", http.NotFoundHandler())
	http.HandleFunc("/posted", h.Posted)
	http.Handle("/posted/", http.NotFoundHandler())
	http.HandleFunc("/dashboard", h.Dashboard)
	http.Handle("/dashboard/", http.NotFoundHandler())
	http.HandleFunc("/profile", h.Profile)
	http.Handle("/profile/", http.NotFoundHandler())

	//Load static folder # Front end
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	//Run and listen to the server
	fmt.Println("listening on: http://localhost:2030/home")
	errServe := http.ListenAndServe(":2030", nil)

	if errServe != nil {
		data := []byte(errServe.Error())
		ioutil.WriteFile("test.txt", data, 0644)
	}
}
