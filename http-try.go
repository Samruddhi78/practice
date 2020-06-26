package main

import (
	"html/template"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w,"<h1><i> Heyyy!! want to take a math quzz?? </i></h1>")
	t, _ := template.ParseFiles("homepage.html")
	st := "first web page"
	t.Execute(w, st)
}
func main() {

	http.HandleFunc("/", HomePageHandler)
	//http.HandleFunc("/quiz", quizPage)
	//to run csss
	http.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("style"))))
	http.ListenAndServe(":8000", nil)
}
