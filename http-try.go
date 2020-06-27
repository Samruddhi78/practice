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
	st := "first web page"
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

/*package main
 2
 3 import (
 4 	"fmt"
 5 	"log"
 6 	"net/http"
 7 	"os"
 8 )
 9
10 func main() {
11 	http.HandleFunc("/", handler)
12 	fmt.Println("listening...")
13 	err := http.ListenAndServe(GetPort(), nil)
14 	if err != nil {
15 		log.Fatal("ListenAndServe: ", err)
16 	}
17 }
18
19 func handler(w http.ResponseWriter, r *http.Request) {
20 	fmt.Fprintf(w, "Hello. This is our first Go web app on Heroku!")
21 }
22
23 // Get the Port from the environment so we can run on Heroku
24 func GetPort() string {
25 	var port = os.Getenv("PORT")
26 	// Set a default port if there is nothing in the environment
27 	if port == "" {
28 		port = "4747"
29 		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
30 	}
31 	return ":" + port
32 }*/
