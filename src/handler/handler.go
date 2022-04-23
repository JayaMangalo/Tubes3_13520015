package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("Path not found")
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))

	// masih dummy data
	data := map[string]string{
		"NamaOrang":    "Andi",
		"NamaPenyakit": "Covid",
		"Status":       "Positif",
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "Ada kesalahan teknis", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Ada kesalahan teknis", http.StatusInternalServerError)
		return
	}
}

// mengambil query string
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Product page : %d", idNumb)
}
