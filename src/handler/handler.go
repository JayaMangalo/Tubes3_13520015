package handler

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"main/src/connector"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func GetInputPenyakitHandler(w http.ResponseWriter, r *http.Request) {

	method := r.Method

	if method == "GET" {
		log.Println("Get ")
		tmpl, err := template.ParseFiles(path.Join("views", "inputPenyakit.html"))
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

}

func PostPenyakitHandler(w http.ResponseWriter, r *http.Request) {

	method := r.Method

	if method == "POST" {
		log.Println("POST")

		r.ParseMultipartForm(10 << 20)
		//namaPenyakit := r.Form.Get("namaPenyakit")
		uploadedFile, handler, err := r.FormFile("file")

		filename := handler.Filename
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer uploadedFile.Close()

		dir, err := os.Getwd()
		fileLocation := filepath.Join(filepath.Dir(dir), "test", "disease")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// NOTE: namafilenya ada tambahan angka random (emg bawaan fungsi dan gabisa diubah)
		tempFile, err := ioutil.TempFile(fileLocation, fmt.Sprintf("%s*.txt", fileNameWithoutExtSliceNotation(filename)))
		if err != nil {
			log.Fatal(err)
		}

		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(uploadedFile)

		if err != nil {
			log.Fatal(err)
		}

		tempFile.Write(fileBytes)

		http.Redirect(w, r, "/", http.StatusFound)
		//fmt.Fprintf(w, "Sucessfully uploaded file")
	}

}

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
func DiagnosisHandler(w http.ResponseWriter, r *http.Request) {
	data := connector.GetDataOrang()
	tmpl, err := template.ParseFiles(path.Join("views", "tes.html"))
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

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
