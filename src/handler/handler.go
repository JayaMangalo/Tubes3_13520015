package handler

import (
	"log"
	"main/src/algorithm"
	"main/src/connector"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func TesDNA(c *gin.Context) {
	nama_penyakit := c.Param("nama_penyakit")
	nama_pengguna := c.Param("nama_pengguna")
	dna_squence := c.Param("dna_sequence")
	currentTime := time.Now()

	date := currentTime.Format("02 April 2006")
	log.Println(date)

	hasil_diagnosis := "False"
	hasil_LCS := algorithm.LCSPercentage(dna_squence, connector.GetSequencePenyakit(nama_penyakit))

	if hasil_LCS >= 80 {
		hasil_diagnosis = "True"
	}
	/*data := gin.H{"Tanggal": date,
	"Nama":      nama_pengguna,
	"Penyakit":  nama_penyakit,
	"diagnosis": hasil_diagnosis}*/

	connector.InsertDataDiagnosis(nama_penyakit, nama_pengguna, hasil_diagnosis, date)
	//c.IndentedJSON(http.StatusOK, data)

}

func PostPenyakitHandler(c *gin.Context) {

	nama_penyakit := c.Param("nama_penyakit")
	dna_squence := c.Param("dna_sequence")

	log.Println(nama_penyakit)
	log.Println(dna_squence)

	if algorithm.IsSanitizedRegex(dna_squence) {
		connector.InsertDataPenyakit(nama_penyakit, dna_squence)
		return
	}

	log.Println("Dna Squence is not correct")
	//fmt.Fprintf(w, "Sucessfully uploaded file")

}

// ngambil data semua orang
func DiagnosisHandler(c *gin.Context) {
	regex := c.Param("query")
	log.Println(regex)
	data := connector.GetDataOrang(regex)

	c.IndentedJSON(http.StatusOK, data)
}

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
