package helpers

import (
	"html/template"
	"net/http"
	"io/ioutil"
	"log"
	"os"
)

func Home(w http.ResponseWriter, req *http.Request) {
	tHome, err := template.ParseFiles("templates/index.html")
	if err != nil {
		data := []byte(err.Error())                               // Permet d'afficher dans le fichier "test.txt"
		ioutil.WriteFile("test.txt", data, 0644)                      // les messages d'erreurs et leurs détails.              
		http.Error(w, err.Error(), http.StatusInternalServerError)        // Affiche  l'utilisateur une erreur 500 en cas d'erreur interne du site
		log.Fatal(err)          
	}

	tHome.Execute(w, nil)
}

func HomeLogged(w http.ResponseWriter, req *http.Request) {
	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	tHomeLogged.Execute(w, nil)
}

func Login(w http.ResponseWriter, req *http.Request) {
	tLogin, err := template.ParseFiles("templates/login.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tLogin.Execute(w, nil)
}

func Register(w http.ResponseWriter, req *http.Request) {
	tRegister, err := template.ParseFiles("templates/register.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tRegister.Execute(w, nil)
}

func Liked(w http.ResponseWriter, req *http.Request) {
	tLiked, err := template.ParseFiles("templates/test4.html")

	if err != nil {
		data := []byte(err.Error())
		var test string
		var test2 string
		var test3 []byte
		test = string(data)
		test2 = "\n" + test
		test3 = []byte(test2)
		ioutil.WriteFile("test.txt", test3, 0644)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}

	tLiked.Execute(w, nil)
}

func Posted(w http.ResponseWriter, req *http.Request) {
	tPosted, err := template.ParseFiles("templates/test10.html")
	if err != nil {
		file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		if _, err := file.WriteString("second line"); err != nil {
			log.Fatal(err)
		}
		
		//Print the contents of the file
		err := ioutil.ReadFile("test.txt")
		if err != nil {
			log.Fatal(err)
		}
	
		test := []byte(err.Error())
		ioutil.WriteFile("test.txt", test, 0644)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		
	}

	tPosted.Execute(w, nil)
}

func Dashboard(w http.ResponseWriter, req *http.Request) {
	tDashboard, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tDashboard.Execute(w, nil)
}

func Profile(w http.ResponseWriter, req *http.Request) {
	tProfile, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tProfile.Execute(w, nil)
}