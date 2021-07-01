package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	src "./src"
	"os"
	//_ "github.com/mattn/go-sqlite3"
)


func main() {

const (
	ConstLogFilename       string        = "test.txt"
)


	logFile, err := os.OpenFile(ConstLogFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)  //ouvre le fichier test.txt
	if err != nil {
		fmt.Printf("[ERROR] - Error opening file: %v", err)
		os.Exit(6)    //ErrorOpenFile
	}
	defer logFile.Close()
	src.Logger = log.New(logFile, "ERROR - ", log.LstdFlags)

	//db, err := sql.Open("sqlite3", "./")
	//Load the pages
	http.Handle("/", http.NotFoundHandler())
	http.HandleFunc("/home", src.Home)
	http.Handle("/home/", http.NotFoundHandler())
	http.HandleFunc("/homeLogged", src.HomeLogged)
	http.Handle("/homeLogged/", http.NotFoundHandler())
	http.HandleFunc("/login", src.Login)
	http.Handle("login/", http.NotFoundHandler())
	http.HandleFunc("/register", src.Register)
	http.Handle("/register/", http.NotFoundHandler())
	http.HandleFunc("/liked", src.Liked)
	http.Handle("/liked/", http.NotFoundHandler())
	http.HandleFunc("/posted", src.Posted)
	http.Handle("/posted/", http.NotFoundHandler())
	http.HandleFunc("/dashboard", src.Dashboard)
	http.Handle("/dashboard/", http.NotFoundHandler())
	http.HandleFunc("/profile", src.Profile)
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