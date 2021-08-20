package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":"+port, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		http.Redirect(w, r, "/success.html", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/#contact", http.StatusSeeOther)
	}
}
