package controllers

import (
	"fmt"
	"rekomendasi_program_studi/config"
	"rekomendasi_program_studi/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler untuk GET semua data program studi
func GetAllProgramStudiHTML(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Get("flash")
	session.Delete("flash")
	session.Save()

	q := c.Query("q")
	var prodi []models.ProgramStudi

	if q != "" {
		config.DB.Where("nama_prodi LIKE ? OR kode LIKE ?", "%"+q+"%", "%"+q+"%").Find(&prodi)
	} else {
		config.DB.Find(&prodi)
	}

	c.HTML(200, "prodi.html", gin.H{
		"Page":  "prodi",
		"Data":  prodi,
		"Nilai": []int{1, 2, 3, 4, 5},
		"Flash": flash,
		"Query": q,
	})
}

// POST /prodi/tambah
func SimpanProgramStudi(c *gin.Context) {
	var prodi models.ProgramStudi
	session := sessions.Default(c)

	prodi.KodeProdi = c.PostForm("kode")
	prodi.NamaProdi = c.PostForm("nama_prodi")
	fmt.Sscanf(c.PostForm("N1"), "%f", &prodi.N1)
	fmt.Sscanf(c.PostForm("N2"), "%f", &prodi.N2)
	fmt.Sscanf(c.PostForm("N3"), "%f", &prodi.N3)
	fmt.Sscanf(c.PostForm("N4"), "%f", &prodi.N4)
	fmt.Sscanf(c.PostForm("N5"), "%f", &prodi.N5)

	if err := config.DB.Create(&prodi).Error; err != nil {
		session.Set("flash", "Gagal menambah prodi")
		session.Save()
		c.Redirect(302, "/prodi")
		return
	}

	session.Set("flash", "Data program studi berhasil ditambahkan")
	session.Save()
	c.Redirect(302, "/prodi")
}

// POST /prodi/edit/:kode
func EditProgramStudi(c *gin.Context) {
	kode := c.Param("kode")
	session := sessions.Default(c)

	var prodi models.ProgramStudi
	if err := config.DB.Where("kode = ?", kode).First(&prodi).Error; err != nil {
		session.Set("flash", "Prodi tidak ditemukan")
		session.Save()
		c.Redirect(302, "/prodi")
		return
	}

	prodi.NamaProdi = c.PostForm("nama_prodi")
	fmt.Sscanf(c.PostForm("N1"), "%f", &prodi.N1)
	fmt.Sscanf(c.PostForm("N2"), "%f", &prodi.N2)
	fmt.Sscanf(c.PostForm("N3"), "%f", &prodi.N3)
	fmt.Sscanf(c.PostForm("N4"), "%f", &prodi.N4)
	fmt.Sscanf(c.PostForm("N5"), "%f", &prodi.N5)

	if err := config.DB.Save(&prodi).Error; err != nil {
		session.Set("flash", "Gagal mengedit prodi")
		session.Save()
		c.Redirect(302, "/prodi")
		return
	}

	session.Set("flash", "Data program studi berhasil diubah")
	session.Save()
	c.Redirect(302, "/prodi")
}

// POST /prodi/hapus/:kode
func HapusProgramStudi(c *gin.Context) {
	kode := c.Param("kode")
	session := sessions.Default(c)

	if err := config.DB.Where("kode = ?", kode).Delete(&models.ProgramStudi{}).Error; err != nil {
		session.Set("flash", "Gagal menghapus prodi")
		session.Save()
		c.Redirect(302, "/prodi")
		return
	}

	session.Set("flash", "Data program studi berhasil dihapus")
	session.Save()
	c.Redirect(302, "/prodi")
}
