package controllers

import (
	"rekomendasi_program_studi/config"
	"rekomendasi_program_studi/models"

	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {
	var jumlahSiswa, jumlahProdi, jumlahIPA, jumlahIPS int64
	var rekap []models.RekapSiswa
	var totalSemua, totalIPA, totalIPS int

	config.DB.Table("data_siswa").Count(&jumlahSiswa)
	config.DB.Table("data_siswa").Where("jurusan = ?", "IPA").Count(&jumlahIPA)
	config.DB.Table("data_siswa").Where("jurusan = ?", "IPS").Count(&jumlahIPS)
	config.DB.Table("program_studi").Count(&jumlahProdi)

	// Ambil rekap seluruh siswa per kelas
	config.DB.Find(&rekap)
	for _, r := range rekap {
		totalSemua += r.JumlahSiswa
		totalIPA += r.IPA
		totalIPS += r.IPS
	}

	c.HTML(200, "dashboard.html", gin.H{
		"Page":        "dashboard",
		"JumlahSiswa": jumlahSiswa,
		"JumlahProdi": jumlahProdi,
		"JumlahIPA":   jumlahIPA,
		"JumlahIPS":   jumlahIPS,
		"Rekap":       rekap,
		"TotalAll":    totalSemua,
		"TotalIPA":    totalIPA,
		"TotalIPS":    totalIPS,
	})
}
