package helpers

import (
	"html/template"
	"net/http"
	"log"

)

var (
	Logger  *log.Logger
)

func Home(w http.ResponseWriter, req *http.Request) {
	tHome, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Home\n", err.Error()) 
			return 
	}

	tHome.Execute(w, nil)
}

func HomeLogged(w http.ResponseWriter, req *http.Request) {
	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: HomeLogged\n", err.Error()) 
			return 
	}

	tHomeLogged.Execute(w, nil)
}

func Login(w http.ResponseWriter, req *http.Request) {
	tLogin, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Login\n", err.Error()) 
			return 
	}

	tLogin.Execute(w, nil)
}

func Register(w http.ResponseWriter, req *http.Request) {
	tRegister, err := template.ParseFiles("templates/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Register\n", err.Error()) 
			return 
	}

	tRegister.Execute(w, nil)
}

func Liked(w http.ResponseWriter, req *http.Request) {
	tLiked, err := template.ParseFiles("templates/liked.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Liked\n", err.Error()) 
			return 
	}

	tLiked.Execute(w, nil)
}

func Posted(w http.ResponseWriter, req *http.Request) {
	tPosted, err := template.ParseFiles("templates/posted.html")
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	Logger.Printf("Erreur interne du serveur sur la page: Posted | ", err.Error()) 
		return 
	}
	tPosted.Execute(w, nil)
}

func Dashboard(w http.ResponseWriter, req *http.Request) {
	tDashboard, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Dashboard\n", err.Error()) 
			return 
	}

	tDashboard.Execute(w, nil)
}

func Profile(w http.ResponseWriter, req *http.Request) {
	tProfile, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Profile\n", err.Error()) 
			return 
	}

	tProfile.Execute(w, nil)
}