package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w,"<h1><i> Heyyy!! want to take a math quzz?? </i></h1>")
	t, _ := template.ParseFiles("homepage.html")
	st := "first web page ,second commit"
	t.Execute(w, st)
}
func main() {

	http.HandleFunc("/", HomePageHandler)
	//http.HandleFunc("/quiz", quizPage)
	//to run csss
	http.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("style"))))
	http.ListenAndServe(GetPort(), nil)
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

//dlhsodfu
