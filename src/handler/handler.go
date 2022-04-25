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

	date := currentTime.Format("2006-01-02 Mon")
	log.Println(date)

	data := gin.H{"Tanggal": date,
		"Nama":      nama_pengguna,
		"Penyakit":  nama_penyakit,
		"diagnosis": algorithm.LCS(dna_squence, connector.GetSequencePenyakit(nama_penyakit))}

	c.IndentedJSON(http.StatusOK, data)

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
	data := connector.GetDataOrang()

	c.IndentedJSON(http.StatusOK, data)
}

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
