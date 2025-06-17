package controllers

import (
	"fmt"
	"net/http"
	"rekomendasi_program_studi/config"
	"rekomendasi_program_studi/models"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetAllSiswaHTML(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Get("flash")
	session.Delete("flash")
	session.Save()

	q := c.Query("q")
	var siswa []models.Siswa

	if q != "" {
		config.DB.Where("nama_siswa LIKE ? OR CAST(nisn AS CHAR) LIKE ?", "%"+q+"%", "%"+q+"%").Find(&siswa)
	} else {
		config.DB.Find(&siswa)
	}

	c.HTML(http.StatusOK, "siswa.html", gin.H{
		"Page":  "siswa",
		"Data":  siswa,
		"Nilai": []int{1, 2, 3, 4, 5},
		"Flash": flash,
		"Query": q,
	})
}

func SimpanSiswa(c *gin.Context) {
	var siswa models.Siswa

	// Konversi NISN dari string ke uint
	nisnStr := c.PostForm("nisn")
	nisn, err := strconv.ParseUint(nisnStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "NISN tidak valid")
		return
	}
	siswa.NISN = uint(nisn)

	siswa.NamaSiswa = c.PostForm("nama_siswa")
	siswa.Jurusan = c.PostForm("jurusan")
	fmt.Sscanf(c.PostForm("N1"), "%f", &siswa.N1)
	fmt.Sscanf(c.PostForm("N2"), "%f", &siswa.N2)
	fmt.Sscanf(c.PostForm("N3"), "%f", &siswa.N3)
	fmt.Sscanf(c.PostForm("N4"), "%f", &siswa.N4)
	fmt.Sscanf(c.PostForm("N5"), "%f", &siswa.N5)

	if err := config.DB.Create(&siswa).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal menyimpan data")
		return
	}

	c.Redirect(http.StatusFound, "/siswa")
}

func EditSiswa(c *gin.Context) {
	nisnStr := c.Param("nisn")
	nisn, err := strconv.ParseUint(nisnStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "NISN tidak valid")
		return
	}

	var siswa models.Siswa
	if err := config.DB.First(&siswa, "nisn = ?", nisn).Error; err != nil {
		c.String(http.StatusNotFound, "Siswa tidak ditemukan")
		return
	}

	// NISN tidak diubah karena merupakan primary key
	siswa.NamaSiswa = c.PostForm("nama_siswa")
	siswa.Jurusan = c.PostForm("jurusan")
	fmt.Sscanf(c.PostForm("N1"), "%f", &siswa.N1)
	fmt.Sscanf(c.PostForm("N2"), "%f", &siswa.N2)
	fmt.Sscanf(c.PostForm("N3"), "%f", &siswa.N3)
	fmt.Sscanf(c.PostForm("N4"), "%f", &siswa.N4)
	fmt.Sscanf(c.PostForm("N5"), "%f", &siswa.N5)

	if err := config.DB.Save(&siswa).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal menyimpan perubahan")
		return
	}

	c.Redirect(http.StatusFound, "/siswa")
}

func HapusSiswa(c *gin.Context) {
	nisnStr := c.Param("nisn")
	nisn, err := strconv.ParseUint(nisnStr, 10, 64)
	session := sessions.Default(c)

	if err != nil {
		session.Set("flash", "NISN tidak valid")
		session.Save()
		c.Redirect(http.StatusFound, "/siswa")
		return
	}

	if err := config.DB.Delete(&models.Siswa{}, "nisn = ?", nisn).Error; err != nil {
		session.Set("flash", "Gagal menghapus data siswa")
	} else {
		session.Set("flash", "Data siswa berhasil dihapus")
	}
	session.Save()
	c.Redirect(http.StatusFound, "/siswa")
}
